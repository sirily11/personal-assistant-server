package user

import (
	"context"
	"fmt"
	"github.com/google/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"personal-assistant/internal/config/constants/keys"
	dto "personal-assistant/internal/dto/authentication"
)

type User[T dto.SignUpUserDtoCredentialData] struct {
	// Unique identifier for the user, following the ObjectID format.
	Id *string `json:"id,omitempty"`

	// Source of the registration
	Source dto.Source `json:"source,omitempty"`

	// List of permissions for the user
	Permission []string `json:"permission,omitempty"`

	CredentialData T `json:"credentialData,omitempty"`
}

type UserRepositoryInterface interface {
	// Create a new user
	Create(user dto.SignUpUserDto[any]) (User[any], error)
}

// UserRepository is the implementation of the UserRepository interface
type UserRepository struct {
	userCollection *mongo.Collection
}

// NewUserRepository creates a new UserRepository
func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		userCollection: db.Collection(keys.UserCollectionKey),
	}
}

// Create a new user
func (r *UserRepository) Create(user dto.SignUpUserDto[any]) (*User[any], error) {
	createdUser := &User[any]{
		Id:         nil,
		Source:     user.Source,
		Permission: user.Permissions,
	}

	// if user uses username and password to register
	// hash the password before saving it to the database
	if c, ok := user.CredentialData.(dto.SignUpUserDtoUserNamePasswordCredentialData); ok {
		u, err := r.createWithUsernameAndPassword(c, createdUser)
		if err != nil {
			return u, err
		}
	}

	// Save the user to the database
	result, err := r.userCollection.InsertOne(context.TODO(), createdUser)
	if err != nil {
		logger.Error("Error saving user to database", err)
		return nil, err
	}

	id := fmt.Sprintf("%s", result.InsertedID)
	createdUser.Id = &id
	return createdUser, nil
}

// createWithUsernameAndPassword creates a new user with username and password
func (r *UserRepository) createWithUsernameAndPassword(c dto.SignUpUserDtoUserNamePasswordCredentialData, createdUser *User[any]) (*User[any], error) {
	hashedPassword, err := r.hashPassword(c.Password)
	if err != nil {
		logger.Error("Error hashing password", err)
		return nil, err
	}

	createdUser.CredentialData = dto.SignUpUserDtoUserNamePasswordCredentialData{
		Type:     c.Type,
		Username: c.Username,
		Password: hashedPassword,
	}
	return nil, nil
}

// hashPassword hashes the password using bcrypt
func (r *UserRepository) hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
