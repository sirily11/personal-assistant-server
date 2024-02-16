package services

import (
	. "sme-demo/internal/repositories/model/whisper"
	"sme-demo/pkgs"
)

type WhisperModelServiceInterface interface {
	Create(whisper WhisperModel) (*WhisperModel, error)
	List() ([]WhisperModel, error)
	GetById(id string) (*WhisperModel, error)
	Delete(id string) error
	Update(id string, whisper WhisperModel) (*WhisperModel, error)
}

type WhisperModelService struct {
	whisperRepository WhisperRepositoryInterface
	s3PreSigner       pkgs.S3PreSignerInterface
}

func NewWhisperModelService(repo WhisperRepositoryInterface, s3PreSigner pkgs.S3PreSignerInterface) WhisperModelServiceInterface {
	return &WhisperModelService{
		whisperRepository: repo,
		s3PreSigner:       s3PreSigner,
	}
}

func (s *WhisperModelService) Create(whisper WhisperModel) (*WhisperModel, error) {
	for i, file := range whisper.FileUrl {
		// generate pre signed url only if url is empty
		if len(file.Url) != 0 {
			continue
		}
		preSignedUrl, err := s.s3PreSigner.GetPreSignedURLForUpload(file.Name)
		if err != nil {
			return nil, err
		}

		whisper.FileUrl[i].Url = preSignedUrl
	}

	return s.whisperRepository.Create(whisper)
}

func (s *WhisperModelService) List() ([]WhisperModel, error) {
	return s.whisperRepository.List()
}

func (s *WhisperModelService) GetById(id string) (*WhisperModel, error) {
	return s.whisperRepository.GetById(id)
}

func (s *WhisperModelService) Delete(id string) error {
	return s.whisperRepository.Delete(id)
}

func (s *WhisperModelService) Update(id string, whisper WhisperModel) (*WhisperModel, error) {
	return s.whisperRepository.Update(id, whisper)
}
