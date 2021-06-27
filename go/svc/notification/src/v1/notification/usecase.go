package notification

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/ispec-inc/monorepo/go/pkg/applog"
	"github.com/ispec-inc/monorepo/go/svc/notification/pkg/registry"
)

type Usecase struct {
	log applog.AppLog
}

func NewUsecase(rgst registry.Registry) Usecase {
	return Usecase{
		log: applog.New(rgst.Logger().New()),
	}
}

func (u Usecase) Send(ctx context.Context, ipt *Input) error {
	spew.Dump(ipt)
	return nil
}
