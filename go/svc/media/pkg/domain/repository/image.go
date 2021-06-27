package repository

import (
	"github.com/ispec-inc/monorepo/go/svc/media/pkg/domain/value"
)

type Image interface {
	Create(localPath, remotePath string, mediaType value.MediaType) (url string, err error)
}
