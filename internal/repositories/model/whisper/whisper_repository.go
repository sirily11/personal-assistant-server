package whisper_repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"sme-demo/internal/config/constants/keys"
)

type WhisperModel struct {
	Id      string      `json:"id"`
	Name    string      `json:"name"`
	Hash    string      `json:"hash"`
	FileUrl []ModelFile `json:"fileUrl"`
	Size    float32     `json:"size"`
}

type ModelFile struct {
	Url           string `json:"url"`
	Name          string `json:"name"`
	IsCoreMLModel bool   `json:"isCoreMLModel"`
}

type WhisperRepositoryInterface interface {
	// Create a new whisper
	Create(whisper WhisperModel) (*WhisperModel, error)
	List() ([]WhisperModel, error)
	GetById(id string) (*WhisperModel, error)
	Delete(id string) error
	Update(id string, whisper WhisperModel) (*WhisperModel, error)
}

type WhisperRepository struct {
	whisperCollection *mongo.Collection
}

func NewWhisperRepository(db *mongo.Database) *WhisperRepository {
	return &WhisperRepository{
		whisperCollection: db.Collection(keys.WhisperModelKey),
	}
}

func (r *WhisperRepository) Create(whisper WhisperModel) (*WhisperModel, error) {
	createdResult, err := r.whisperCollection.InsertOne(nil, whisper)
	if err != nil {
		return nil, err
	}

	whisper.Id = createdResult.InsertedID.(string)
	return &whisper, nil
}

func (r *WhisperRepository) List() ([]WhisperModel, error) {
	var whispers []WhisperModel
	cursor, err := r.whisperCollection.Find(nil, nil)
	if err != nil {
		return nil, err
	}

	err = cursor.All(nil, &whispers)
	if err != nil {
		return nil, err
	}

	return whispers, nil
}

func (r *WhisperRepository) GetById(id string) (*WhisperModel, error) {
	var whisper WhisperModel
	err := r.whisperCollection.FindOne(nil, nil).Decode(&whisper)
	if err != nil {
		return nil, err
	}

	return &whisper, nil
}

func (r *WhisperRepository) Delete(id string) error {
	_, err := r.whisperCollection.DeleteOne(nil, nil)
	return err
}

func (r *WhisperRepository) Update(id string, whisper WhisperModel) (*WhisperModel, error) {
	_, err := r.whisperCollection.ReplaceOne(nil, nil, nil)
	if err != nil {
		return nil, err
	}

	return &whisper, nil
}
