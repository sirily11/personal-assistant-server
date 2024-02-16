package main

import (
	"io"
	"os/exec"

	"github.com/google/logger"
)

func main() {
	logger.Init("Logger", true, false, io.Discard)

	output, err := exec.Command(
		"openapi-generator",
		"generate",
		"-i",
		"./api/openapi_schema_v2_2024_02_15.yaml",
		"-g",
		"go-gin-server",
		"-o",
		"./internal/api",
		"-p",
		"packageName=api",
	).CombinedOutput()
	logger.Info("Running command: openapi-generator generate -i ../api/openapi_schema_v2_2024_02_15.yaml -g go-gin-server -o ./api -p packageName=api")
	if err != nil {
		logger.Error(err)
		logger.Fatal(string(output))
	}
	logger.Info(string(output))
}
