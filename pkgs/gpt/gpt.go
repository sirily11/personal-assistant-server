package gpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sme-demo/internal/config"
	"sme-demo/internal/config/constants/environments"
	"sme-demo/internal/dto/gpt"
	"sme-demo/pkgs/gpt/functions"
	"sme-demo/pkgs/template"

	"github.com/google/logger"
)

type IGptClient interface {
	Generate(prompt *string, history []gpt.MessageResponseDto) (*string, []gpt.MessageResponseDto, error)
	SetFunctionStore(store functions.FunctionStore)
	SetConfig(config config.Config)
}

type Client struct {
	functions []functions.IFunction
	store     functions.FunctionStore
	config    config.Config
	template  *template.TemplateClient
}

// NewGptClient returns a new instance of GptClient.
func NewGptClient(template *template.TemplateClient) IGptClient {
	return &Client{
		functions: []functions.IFunction{},
		template:  template,
	}
}

func (g *Client) SetFunctionStore(store functions.FunctionStore) {
	g.store = store
}

func (g *Client) SetConfig(config config.Config) {
	g.config = config
}

// Generate takes in a prompt and returns a generated errors.
func (g *Client) Generate(prompt *string, history []gpt.MessageResponseDto) (*string, []gpt.MessageResponseDto, error) {
	messages := g.createMessages(prompt, history)
	logger.Infof("<GPTRequest messages=%v/>", messages)

	body := gpt.RequestDto{
		Messages:  messages,
		Functions: g.generateFunctions(),
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, nil, err
	}
	request, err := http.NewRequest(http.MethodPost, environments.GptEndpoint, bytes.NewBuffer(jsonBody))
	request.Header.Add("api-key", environments.GptKey)

	if err != nil {
		return nil, nil, err
	}

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return nil, nil, err
	}

	// log errors body
	//var result map[string]interface{}
	var result gpt.ResponseDto
	var stringBody string

	// read errors body
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	stringBody = buf.String()

	// write back to errors body
	response.Body = io.NopCloser(bytes.NewBufferString(stringBody))

	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return nil, nil, err
	}

	// if gpt api throws an error, return empty string
	if response.StatusCode != http.StatusOK {
		logger.Errorf("GPT request failed with status code: %v, %s", response.StatusCode, stringBody)
		return nil, nil, fmt.Errorf("gpt request failed with status code: %v", response.StatusCode)
	}

	// use function if there is one
	message := result.Choices[0].Message
	logger.Infof(stringBody)
	logger.Infof("<GPTResponse content=%s function=%v name=%v role=%v />", message.Content, message.FunctionCall, message.Name, message.Role)

	newHistory := append(messages, gpt.MessageResponseDto{
		Role:         message.Role,
		Content:      message.Content,
		Name:         message.Name,
		FunctionCall: message.FunctionCall,
	})
	newHistory, err = g.useFunction(message, newHistory)
	if err != nil {
		logger.Error(err)
		return nil, nil, err
	}

	defer response.Body.Close()

	content := newHistory[len(newHistory)-1].Content
	return &content, newHistory, nil
}

func (g *Client) generateFunctions() []gpt.FunctionResponseDto {
	var returnedFunctions []gpt.FunctionResponseDto
	for _, function := range g.functions {
		returnedFunctions = append(returnedFunctions, gpt.FunctionResponseDto{
			Name:        function.Name(),
			Description: function.Description(),
			Parameters:  function.Parameters(),
		})
	}

	return returnedFunctions
}

func (g *Client) useFunction(result gpt.MessageResponseDto, history []gpt.MessageResponseDto) ([]gpt.MessageResponseDto, error) {
	newHistory := history
	if result.FunctionCall != nil {
		for _, function := range g.functions {
			if function.Name() == result.FunctionCall.Name {
				var functionArguments map[string]interface{}
				err := json.Unmarshal([]byte(result.FunctionCall.Arguments), &functionArguments)
				if err != nil {
					return nil, err
				}
				function.SetStore(g.store)
				result, err := function.Execute(functionArguments)
				if err != nil {
					return newHistory, err
				}
				logger.Infof("Function %v executed with result %v", function.Name(), result)

				// Add history
				functionName := function.Name()
				newHistory = append(newHistory, gpt.MessageResponseDto{
					Role:    RoleFunction,
					Content: result.Content,
					Name:    &functionName,
				})
				if function.Config().UseGptToInterpretResponses {
					_, newHistory, err := g.Generate(nil, newHistory)
					if err != nil {
						return nil, err
					}
					return newHistory, nil
				}

				return newHistory, nil
			}
		}
		return nil, fmt.Errorf("function %v not found", result.FunctionCall.Name)
	}

	return newHistory, nil
}

// createMessages creates a list of messages with history and prompt included.
func (g *Client) createMessages(prompt *string, history []gpt.MessageResponseDto) []gpt.MessageResponseDto {
	var messages []gpt.MessageResponseDto
	// only add system message if there is no history
	if len(history) == 0 {
		renderedPrompt, err := g.template.Render(g.config.Prompt.Prompt)
		if err != nil {
			logger.Error(err)
			return nil
		}
		messages = append(messages, gpt.MessageResponseDto{
			Role:    "system",
			Content: renderedPrompt,
		})
	}
	for _, message := range history {
		messages = append(messages, message)
	}

	if prompt != nil {
		messages = append(messages, gpt.MessageResponseDto{
			Role:    "user",
			Content: *prompt,
		})
	}
	return messages
}
