package aws

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3UploadInput struct {
	Bucket string
	Region string
	Path   string
	File   string
	ACL    acl
}

func (a awsImpl) S3Upload(ipt S3UploadInput) (url string, aerr Error) {
	f, err := os.Open(ipt.File)
	if err != nil {
		return url, newFileError(err)
	}
	defer os.Remove(ipt.File)

	sess := session.Must(session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(a.accessKey, a.secretKey, ""),
		Region:      aws.String(ipt.Region),
	}))

	uploader := s3manager.NewUploader(sess)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(ipt.Bucket),
		Key:         aws.String(ipt.Path),
		Body:        f,
		ContentType: aws.String(contentType(ipt.File)),
		ACL:         aws.String(string(ipt.ACL)),
	})

	if err == nil {
		url = fmt.Sprintf("https://%s.s3-%s.amazonaws.com/%s", ipt.Bucket, ipt.Region, ipt.Path)
		return url, nil
	}

	if err, ok := err.(awserr.Error); ok && err.Code() == request.CanceledErrorCode {
		return url, newNetworkError(fmt.Errorf("upload canceled due to timeout: %v", err))
	}
	return url, newNetworkError(fmt.Errorf("failed to upload object, %v", err))
}

func contentType(file string) string {
	// Please add cases which you need
	switch filepath.Ext(file) {
	case "jpg":
		return "image/jpeg"
	case "jpeg":
		return "image/jpeg"
	case "gif":
		return "image/gif"
	case "png":
		return "image/png"
	default:
		return "image/png"
	}
}
