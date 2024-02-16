package whisper_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"personal-assistant/internal/config/constants/keys"
)

type WhisperModel struct {
	Id      *string     `json:"id" bson:"_id,omitempty"`
	Name    string      `json:"name" bson:"name" binding:"required"`
	FileUrl []ModelFile `json:"fileUrl" binding:"required" bson:"fileUrl"`
}

type ModelFile struct {
	Url           string `json:"url" bson:"url" binding:"required,url" validate:"url"`
	Name          string `json:"name" binding:"required" bson:"name"`
	IsCoreMLModel bool   `json:"isCoreMLModel" bson:"isCoreMLModel" binding:"required"`
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

func NewWhisperRepository(db *mongo.Database) WhisperRepositoryInterface {
	return &WhisperRepository{
		whisperCollection: db.Collection(keys.WhisperModelKey),
	}
}

// Create a new whisper
func (r *WhisperRepository) Create(whisper WhisperModel) (*WhisperModel, error) {
	createdResult, err := r.whisperCollection.InsertOne(context.TODO(), whisper)
	if err != nil {
		return nil, err
	}

	id := createdResult.InsertedID.(primitive.ObjectID).Hex()
	whisper.Id = &id
	return &whisper, nil
}

// List all whispers
func (r *WhisperRepository) List() ([]WhisperModel, error) {
	var whispers = make([]WhisperModel, 0)
	cursor, err := r.whisperCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	err = cursor.All(nil, &whispers)
	if err != nil {
		return nil, err
	}

	return whispers, nil
}

// GetById returns a whisper by its id
func (r *WhisperRepository) GetById(id string) (*WhisperModel, error) {
	var whisper WhisperModel
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = r.whisperCollection.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&whisper)
	if err != nil {
		return nil, err
	}

	return &whisper, nil
}

// Delete a whisper by its id
func (r *WhisperRepository) Delete(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.whisperCollection.DeleteOne(context.TODO(), bson.M{"_id": objectId})
	return err
}

// Update a whisper by its id
func (r *WhisperRepository) Update(id string, whisper WhisperModel) (*WhisperModel, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	_, err = r.whisperCollection.ReplaceOne(context.TODO(), bson.M{"_id": objectId}, whisper)
	if err != nil {
		return nil, err
	}

	whisper.Id = &id

	return &whisper, nil
}
