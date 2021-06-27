package image

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/ispec-inc/monorepo/go/pkg/presenter"
	"github.com/ispec-inc/monorepo/go/pkg/uuid"
	pb "github.com/ispec-inc/monorepo/go/proto/media/api/rest/v1/image"
	"github.com/ispec-inc/monorepo/go/svc/media/pkg/registry"
	"github.com/ispec-inc/monorepo/go/svc/media/src/v1/image"
)

const (
	MediaTmpDir = "."
)

type handler struct {
	usecase image.Usecase
}

func newHandler(rgst registry.Registry) handler {
	uc := image.NewUsecase(rgst)
	return handler{uc}
}

func (h handler) Create(w http.ResponseWriter, r *http.Request) {
	reqimg, header, _ := r.FormFile("image")
	path := fmt.Sprintf("%s/%s%s", MediaTmpDir, uuid.GenerateTimeUUID(), filepath.Ext(header.Filename))
	img, _ := os.Create(path)
	defer img.Close()
	io.Copy(img, reqimg)

	opt, err := h.usecase.Create(r.Context(), &image.Input{
		Filepath: path,
	})
	if err != nil {
		presenter.ApplicationException(w, err)
		return
	}

	presenter.Response(w, &pb.CreateResponse{
		Url: opt.URL,
	})
}
