package config

const (
	ConditionTrue = "true"
	// ConditionFalse is a constant for false condition in the config
	ConditionFalse = "false"
)

type Config struct {
	Storage Storage `mapstructure:"storage"`
}

type Storage struct {
	S3 S3 `mapstructure:"s3"`
}

type S3 struct {
	// ModelUploadPrefix is the prefix for the model upload path.
	// For example, if the prefix is "models", the model will be uploaded to "s3://bucket/models"
	ModelUploadPrefix string `mapstructure:"modelUploadPath"`
	Region            string `mapstructure:"region"`
	Bucket            string `mapstructure:"bucket"`
}
