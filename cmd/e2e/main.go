package main

import (
	"github.com/google/logger"
	"io"
	"os/exec"
	"path/filepath"
)

func getAllTests(targetDir string) []string {
	// get all files end with .http in the targetDir
	path := filepath.Join(targetDir, "*.http")
	files, err := filepath.Glob(path)
	if err != nil {
		logger.Fatal(err)
	}
	return files
}

func main() {
	logger.Init("Logger", true, false, io.Discard)

	logger.Info("Downloading ijhttp")
	if _, err := exec.LookPath("./ijhttp.zip"); err != nil {
		output, err := exec.Command("curl", "-f", "-L", "-o", "ijhttp.zip", "https://jb.gg/ijhttp/latest").Output()
		if err != nil {
			logger.Fatal(string(output))
		}
	} else {
		logger.Info("ijhttp.zip already exists")
	}

	logger.Info("Setting up ijhttp")
	output, err := exec.Command("chmod", "+x", "ijhttp.zip").Output()
	if err != nil {
		logger.Fatal(string(output))
	}

	logger.Info("Unzipping ijhttp")
	if _, err := exec.LookPath("./ijhttp/ijhttp"); err != nil {
		cmd := exec.Command("unzip", "-q", "ijhttp.zip")
		if err := cmd.Run(); err != nil {
			logger.Fatal(err)
		}
	} else {
		logger.Info("ijhttp already exists")
	}

	files := getAllTests("tests")
	for _, file := range files {
		logger.Info("Running test: ", file)
		output, err := exec.Command("./ijhttp/ijhttp", file, "--env-file", "tests/http-client.env.json", "--env", "dev").Output()
		if err != nil {
			logger.Fatal(string(output))
		}
		logger.Info(string(output))
	}
}
