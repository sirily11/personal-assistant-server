package whisper_repository

import (
	"context"
	"github.com/google/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/tryvium-travels/memongo"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
	"sme-demo/internal/config/constants/keys"
	"sme-demo/internal/repositories"
	"testing"
)

type RepositoryTestSuite struct {
	suite.Suite
	repo   WhisperRepositoryInterface
	db     *mongo.Database
	server *memongo.Server
}

func (suite *RepositoryTestSuite) SetupTest() {
	logger.Init("Logger", true, false, io.Discard)
	mongoServer, err := memongo.Start(keys.TestMongoDBVersion)
	if err != nil {
		suite.T().Fatal(err)
	}

	dbClient := repositories.NewDatabase()
	db := dbClient.ConnectWithUriAndDBName(mongoServer.URI(), keys.TestMongoDatabase)

	suite.repo = NewWhisperRepository(db)
	suite.db = db
	suite.server = mongoServer
}

func (suite *RepositoryTestSuite) TearDownTest() {
	suite.db.Client().Disconnect(context.TODO())
	suite.server.Stop()
}

func (suite *RepositoryTestSuite) TestCreateWhisperModel() {
	whisper := WhisperModel{
		Name: "test",
		FileUrl: []ModelFile{
			{
				Url:           "test",
				Name:          "test",
				IsCoreMLModel: false,
			},
		},
	}
	create, err := suite.repo.Create(whisper)
	if err != nil {
		suite.T().Fatal(err)
	}

	assert.NotNil(suite.T(), create)
	assert.Equal(suite.T(), whisper.Name, create.Name)
	assert.NotEmpty(suite.T(), create.Id)
}

func (suite *RepositoryTestSuite) TestListWhisperModel() {
	whisper := WhisperModel{
		Name: "test",
		FileUrl: []ModelFile{
			{
				Url:           "test",
				Name:          "test",
				IsCoreMLModel: false,
			},
		},
	}
	_, err := suite.repo.Create(whisper)
	if err != nil {
		suite.T().Fatal(err)
	}

	list, err := suite.repo.List()
	if err != nil {
		suite.T().Fatal(err)
	}

	assert.NotNil(suite.T(), list)
	assert.NotEmpty(suite.T(), list)
}

func (suite *RepositoryTestSuite) TestListWhisperModelWithoutData() {
	list, err := suite.repo.List()
	if err != nil {
		suite.T().Fatal(err)
	}

	assert.NotNil(suite.T(), list)
	assert.Empty(suite.T(), list)
}

func (suite *RepositoryTestSuite) TestGetByIdWhisperModel() {
	whisper := WhisperModel{
		Name: "test",
		FileUrl: []ModelFile{
			{
				Url:           "test",
				Name:          "test",
				IsCoreMLModel: false,
			},
		},
	}
	create, err := suite.repo.Create(whisper)
	if err != nil {
		suite.T().Fatal(err)
	}

	get, err := suite.repo.GetById(*create.Id)
	if err != nil {
		suite.T().Fatal(err)
	}

	assert.NotNil(suite.T(), get)
	assert.Equal(suite.T(), whisper.Name, get.Name)
	assert.NotEmpty(suite.T(), get.Id)
}

func (suite *RepositoryTestSuite) TestDeleteWhisperModel() {
	whisper := WhisperModel{
		Name: "test",
		FileUrl: []ModelFile{
			{
				Url:           "test",
				Name:          "test",
				IsCoreMLModel: false,
			},
		},
	}
	create, err := suite.repo.Create(whisper)
	if err != nil {
		suite.T().Fatal(err)
	}

	err = suite.repo.Delete(*create.Id)
	if err != nil {
		suite.T().Fatal(err)
	}

	get, err := suite.repo.GetById(*create.Id)
	assert.Nil(suite.T(), get)
	assert.NotNil(suite.T(), err)
}

func (suite *RepositoryTestSuite) TestUpdateWhisperModel() {
	whisper := WhisperModel{
		Name: "test",
		FileUrl: []ModelFile{
			{
				Url:           "test",
				Name:          "test",
				IsCoreMLModel: false,
			},
		},
	}
	create, err := suite.repo.Create(whisper)
	if err != nil {
		suite.T().Fatal(err)
	}

	whisper.Name = "test2"
	update, err := suite.repo.Update(*create.Id, whisper)
	if err != nil {
		suite.T().Fatal(err)
	}

	assert.NotNil(suite.T(), update)
	assert.Equal(suite.T(), whisper.Name, update.Name)
	assert.NotEmpty(suite.T(), update.Id)
}

func TestRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}
