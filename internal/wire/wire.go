/// This file is used to generate wire.go file using wire tool
/// Run `go generate ./...` to generate wire.go file

//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"personal-assistant/internal/config"
	"personal-assistant/internal/controllers"
	whisper_repository "personal-assistant/internal/repositories/model/whisper"
	"personal-assistant/internal/services"
	"personal-assistant/pkgs"
)

func InitializeWhisperController(config config.Config, db *mongo.Database) *controllers.WhisperController {
	wire.Build(controllers.NewWhisperController, whisper_repository.NewWhisperRepository, services.NewWhisperModelService, pkgs.NewS3PreSigner)
	return &controllers.WhisperController{}
}
