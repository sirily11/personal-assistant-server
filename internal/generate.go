package internal

//go:generate mockgen -source=./repositories/model/whisper/whisper_repository.go -destination=./repositories/model/whisper/mock_whisper_repository.go -package=whisper_repository
//go:generate mockgen -source=./services/whisper_service.go -destination=./services/mock_whisper_service.go -package=services
//go:generate mockgen -source=../pkgs/s3_presigner.go -destination=../pkgs/mock_s3_presigner.go -package=pkgs

//go:generate wire ./wire/wire.go
