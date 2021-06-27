package config

import (
	"github.com/caarlos0/env/v6"
)

func init() {
	if err := env.Parse(&AWS); err != nil {
		panic(err)
	}
}

var AWS aws

type aws struct {
	AccessKey      string `env:"MEDIA_AWS_ACCESS_KEY"`
	SecretKey      string `env:"MEDIA_AWS_SECRET_KEY"`
	S3BucketName   string `env:"MEDIA_AWS_S3_BUCKET_NAME"`
	S3BucketRegion string `env:"MEDIA_AWS_S3_BUCKET_REGION"`
}
