package registry

import (
	"github.com/ispec-inc/monorepo/go/pkg/aws"
	"github.com/ispec-inc/monorepo/go/svc/media/pkg/config"
	"github.com/ispec-inc/monorepo/go/svc/media/pkg/domain/repository"
	"github.com/ispec-inc/monorepo/go/svc/media/pkg/infra/dao"
)

type Repository struct {
	aws aws.AWS
}

func NewRepository() (Repository, func() error, error) {
	r := Repository{
		aws: aws.NewAWS(config.AWS.AccessKey, config.AWS.SecretKey),
	}
	cleanup := func() error { return nil }
	return r, cleanup, nil
}

func (r Repository) NewImage() repository.Image {
	return dao.NewS3(r.aws)
}
