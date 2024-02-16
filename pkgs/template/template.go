package template

import (
	"bytes"
	"html/template"
	"sme-demo/internal/config"
)

type TemplateClient struct {
	config   config.Config
	chatroom ChatRoomData
}

type ChatRoomData struct {
}

func NewTemplateClient(config config.Config) *TemplateClient {
	return &TemplateClient{
		config: config,
	}
}

func (t *TemplateClient) SetChatRoom(chatroom ChatRoomData) {
	t.chatroom = chatroom
}

func (t *TemplateClient) Render(content string) (string, error) {
	tmpl, err := template.New("render").Parse(content)
	if err != nil {
		return "", err
	}

	type Data struct {
		ChatRoom ChatRoomData
		Config   config.Config
	}

	var result bytes.Buffer
	err = tmpl.Execute(&result, Data{
		ChatRoom: t.chatroom,
		Config:   t.config,
	})

	if err != nil {
		return "", err
	}

	return result.String(), nil
}
