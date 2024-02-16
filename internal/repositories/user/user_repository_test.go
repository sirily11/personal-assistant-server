package user

import (
	"context"
	"github.com/google/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/tryvium-travels/memongo"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
	"sme-demo/internal/config/constants/keys"
	dto "sme-demo/internal/dto/authentication"
	"sme-demo/internal/repositories"
	"testing"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	repo   *UserRepository
	db     *mongo.Database
	server *memongo.Server
}

func (suite *UserRepositoryTestSuite) SetupTest() {
	logger.Init("Logger", true, false, io.Discard)
	mongoServer, err := memongo.Start(keys.TestMongoDBVersion)
	if err != nil {
		suite.T().Fatal(err)
	}

	dbClient := repositories.NewDatabase()
	db := dbClient.ConnectWithUriAndDBName(mongoServer.URI(), keys.TestMongoDatabase)

	suite.repo = NewUserRepository(db)
	suite.db = db
	suite.server = mongoServer
}

func (suite *UserRepositoryTestSuite) TearDownTest() {
	suite.db.Client().Disconnect(context.TODO())
	suite.server.Stop()
}

func (suite *UserRepositoryTestSuite) TestCreateWithUsernamePassword() {
	user := dto.SignUpUserDto[any]{
		Source:      dto.SourceWebApp,
		Permissions: []string{},
		CredentialData: dto.SignUpUserDtoUserNamePasswordCredentialData{
			Username: "test",
			Password: "test",
		},
	}
	create, err := suite.repo.Create(user)
	if err != nil {
		suite.T().Fatal(err)
	}

	assert.NotNil(suite.T(), create)
	assert.Equal(suite.T(), user.Source, create.Source)
	assert.NotEmpty(suite.T(), create.Id)
	assert.NotEqual(suite.T(), create.CredentialData.(dto.SignUpUserDtoUserNamePasswordCredentialData).Password, "test")
	assert.Greater(suite.T(), len(create.CredentialData.(dto.SignUpUserDtoUserNamePasswordCredentialData).Password), 4)
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
