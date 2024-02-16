/// This file is used to generate wire.go file using wire tool
/// Run `go generate ./...` to generate wire.go file

//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"sme-demo/internal/config"
	"sme-demo/internal/controllers"
	whisper_repository "sme-demo/internal/repositories/model/whisper"
	"sme-demo/internal/services"
	"sme-demo/pkgs"
)

func InitializeWhisperController(config config.Config, db *mongo.Database) *controllers.WhisperController {
	wire.Build(controllers.NewWhisperController, whisper_repository.NewWhisperRepository, services.NewWhisperModelService, pkgs.NewS3PreSigner)
	return &controllers.WhisperController{}
}
