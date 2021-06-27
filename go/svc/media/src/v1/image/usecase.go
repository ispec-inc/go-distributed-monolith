package image

import (
	"context"
	"path/filepath"

	"github.com/ispec-inc/monorepo/go/svc/media/pkg/domain/repository"
	"github.com/ispec-inc/monorepo/go/svc/media/pkg/domain/value"
	"github.com/ispec-inc/monorepo/go/svc/media/pkg/registry"
)

type Usecase struct {
	imageRepository repository.Image
}

func NewUsecase(rgst registry.Registry) Usecase {
	return Usecase{
		imageRepository: rgst.Repository().NewImage(),
	}
}

func (u Usecase) Create(ctx context.Context, ipt *Input) (*Output, error) {
	remotePath := filepath.Base(ipt.Filepath)
	url, err := u.imageRepository.Create(ipt.Filepath, remotePath, value.MediaTypeImage)
	return &Output{
		URL: url,
	}, err
}
