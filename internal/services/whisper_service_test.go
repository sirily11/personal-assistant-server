package services

import (
	"context"
	"github.com/google/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/tryvium-travels/memongo"
	"go.uber.org/mock/gomock"
	"io"
	"sme-demo/internal/config/constants/keys"
	"sme-demo/internal/repositories"
	whisper_repository "sme-demo/internal/repositories/model/whisper"
	"sme-demo/pkgs"
	"testing"
)

type TestSuite struct {
	suite.Suite
	ctrl *gomock.Controller
}

func (suite *TestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
}

func (suite *TestSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func (suite *TestSuite) TestCreateWithoutFileToUpload() {
	repo := whisper_repository.NewMockWhisperRepositoryInterface(suite.ctrl)
	repo.EXPECT().Create(gomock.Any()).Return(&whisper_repository.WhisperModel{}, nil)

	signer := pkgs.NewMockS3PreSignerInterface(suite.ctrl)

	service := NewWhisperModelService(repo, signer)
	_, err := service.Create(whisper_repository.WhisperModel{
		FileUrl: []whisper_repository.ModelFile{
			{
				Name: "file",
				Url:  "url",
			},
		},
	})

	assert.Nil(suite.T(), err)
}

func (suite *TestSuite) TestCreateWithFileToUpload() {
	logger.Init("Logger", true, false, io.Discard)
	mongoServer, err := memongo.Start(keys.TestMongoDBVersion)
	if err != nil {
		suite.T().Fatal(err)
	}

	dbClient := repositories.NewDatabase()
	db := dbClient.ConnectWithUriAndDBName(mongoServer.URI(), keys.TestMongoDatabase)

	defer db.Client().Disconnect(context.TODO())

	repo := whisper_repository.NewWhisperRepository(db)

	signer := pkgs.NewMockS3PreSignerInterface(suite.ctrl)
	signer.EXPECT().GetPreSignedURLForUpload("file").Return("url", nil)

	service := NewWhisperModelService(repo, signer)
	data, err := service.Create(whisper_repository.WhisperModel{
		FileUrl: []whisper_repository.ModelFile{
			{
				Name: "file",
			},
		},
	})

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), data)
	assert.Equal(suite.T(), "url", data.FileUrl[0].Url)
}

func (suite *TestSuite) TestList() {
	repo := whisper_repository.NewMockWhisperRepositoryInterface(suite.ctrl)
	repo.EXPECT().List().Return([]whisper_repository.WhisperModel{}, nil)

	signer := pkgs.NewMockS3PreSignerInterface(suite.ctrl)

	service := NewWhisperModelService(repo, signer)
	_, err := service.List()

	assert.Nil(suite.T(), err)
}

func (suite *TestSuite) TestGetById() {
	repo := whisper_repository.NewMockWhisperRepositoryInterface(suite.ctrl)
	repo.EXPECT().GetById("id").Return(&whisper_repository.WhisperModel{}, nil)

	signer := pkgs.NewMockS3PreSignerInterface(suite.ctrl)

	service := NewWhisperModelService(repo, signer)
	_, err := service.GetById("id")

	assert.Nil(suite.T(), err)
}

func (suite *TestSuite) TestDelete() {
	repo := whisper_repository.NewMockWhisperRepositoryInterface(suite.ctrl)
	repo.EXPECT().Delete("id").Return(nil)

	signer := pkgs.NewMockS3PreSignerInterface(suite.ctrl)

	service := NewWhisperModelService(repo, signer)
	err := service.Delete("id")

	assert.Nil(suite.T(), err)
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
