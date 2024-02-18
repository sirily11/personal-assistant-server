package pkgs

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/logger"
	"net/url"
	"personal-assistant/internal/config"
)

type S3PreSignerInterface interface {
	GetPreSignedURLForUpload(name string) (string, error)
}

type S3PreSigner struct {
	PreSigner    *s3.PresignClient
	systemConfig config.Config
}

func NewS3PreSigner(systemConfig config.Config) S3PreSignerInterface {
	cfg, err := awsConfig.LoadDefaultConfig(context.TODO())
	if err != nil {
		logger.Fatal("unable to load SDK config, " + err.Error())
	}

	return &S3PreSigner{
		PreSigner: s3.NewPresignClient(
			s3.NewFromConfig(
				cfg,
			),
		),
		systemConfig: systemConfig,
	}
}

// GetPreSignedURLForUpload returns a preSigned URL for uploading an object to S3
func (s *S3PreSigner) GetPreSignedURLForUpload(name string) (string, error) {
	key, err := url.JoinPath(s.systemConfig.Storage.S3.Whisper.ModelUploadPrefix, name)
	if err != nil {
		return "", err
	}
	obj, err := s.PreSigner.PresignPutObject(
		context.TODO(),
		&s3.PutObjectInput{
			Bucket: aws.String(s.systemConfig.Storage.S3.General.Bucket),
			Key:    aws.String(key),
		},
	)
	if err != nil {
		return "", err
	}

	return obj.URL, nil
}
