package aws

type AWS interface {
	S3Upload(ipt S3UploadInput) (url string, err Error)
}

type awsImpl struct {
	accessKey string
	secretKey string
}

func NewAWS(accessKey, secretKey string) AWS {
	return awsImpl{
		accessKey: accessKey,
		secretKey: secretKey,
	}
}
