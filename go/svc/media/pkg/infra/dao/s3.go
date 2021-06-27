package dao

import (
	"fmt"

	"github.com/ispec-inc/monorepo/go/pkg/apperror"
	"github.com/ispec-inc/monorepo/go/pkg/aws"
	"github.com/ispec-inc/monorepo/go/svc/media/pkg/config"
	"github.com/ispec-inc/monorepo/go/svc/media/pkg/domain/value"
)

type S3 struct {
	aws aws.AWS
}

func NewS3(a aws.AWS) S3 {
	return S3{a}
}

func (d S3) Create(localPath, remotePath string, mtype value.MediaType) (string, error) {
	remotePath = fmt.Sprintf("%s/%s", mtype.String(), remotePath)

	url, err := d.aws.S3Upload(aws.S3UploadInput{
		Bucket: config.AWS.S3BucketName,
		Region: config.AWS.S3BucketRegion,
		Path:   remotePath,
		File:   localPath,
		ACL:    aws.ACLPublicRead,
	})
	if err != nil {
		return "", apperror.NewAWS(err)
	}

	return url, nil
}
