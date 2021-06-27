package article

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/ispec-inc/monorepo/go/pkg/applog"
	"github.com/ispec-inc/monorepo/go/pkg/msgbs"
	"github.com/ispec-inc/monorepo/go/pkg/presenter"
	pb "github.com/ispec-inc/monorepo/go/proto/admin/api/rest/v1/article"
	"github.com/ispec-inc/monorepo/go/svc/admin/pkg/logger"
	"github.com/ispec-inc/monorepo/go/svc/admin/pkg/model"
	"github.com/ispec-inc/monorepo/go/svc/admin/pkg/redisbs"
	"github.com/ispec-inc/monorepo/go/svc/admin/pkg/view"
)

type controller struct {
	log applog.AppLog
	bs  msgbs.MessageBus
}

func newController() controller {
	return controller{
		log: applog.New(logger.Get()),
		bs:  redisbs.Get(),
	}
}

func (c controller) list(w http.ResponseWriter, r *http.Request) {
	atls := &model.Articles{}
	if err := atls.Find(); err != nil {
		c.log.Error(r.Context(), err)
		presenter.ApplicationException(w, err)
		return
	}

	presenter.Response(w, &pb.ListResponse{
		Articles: view.NewArticles(atls),
	})
}

func (c controller) get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		presenter.BadRequestError(w, err)
		return
	}

	atl := &model.Article{}
	if err := atl.Find(int64(id)); err != nil {
		c.log.Error(r.Context(), err)
		presenter.ApplicationException(w, err)
		return
	}

	presenter.Response(w, &pb.GetResponse{
		Article: view.NewArticle(atl),
	})
}

func (c controller) create(w http.ResponseWriter, r *http.Request) {
	body := &pb.CreateRequest{}
	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		presenter.BadRequestError(w, err)
		return
	}
	if err := body.Validate(true); err != nil {
		presenter.BadRequestError(w, err)
		return
	}

	atl := model.NewArticle(body.UserId, body.Title, body.Body)
	if err := atl.Create(); err != nil {
		c.log.Error(r.Context(), err)
		presenter.ApplicationException(w, err)
		return
	}

	atls := &model.Articles{}
	if err := atls.Find(); err != nil {
		c.log.Error(r.Context(), err)
		presenter.ApplicationException(w, err)
		return
	}

	c.bs.Publish(msgbs.AddArticle, msgbs.Article{Title: atl.Title})
	presenter.Response(w, &pb.CreateResponse{
		Articles: view.NewArticles(atls),
	})
}

func (c controller) update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		presenter.BadRequestError(w, err)
		return
	}
	body := &pb.UpdateRequest{}
	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		presenter.BadRequestError(w, err)
		return
	}
	if err := body.Validate(true); err != nil {
		presenter.BadRequestError(w, err)
		return
	}

	atl := &model.Article{}
	if err := atl.Find(int64(id)); err != nil {
		c.log.Error(r.Context(), err)
		presenter.ApplicationException(w, err)
		return
	}
	atl.UserID = body.UserId
	atl.Title = body.Title
	atl.Body = body.Body
	if err := atl.Save(); err != nil {
		c.log.Error(r.Context(), err)
		presenter.ApplicationException(w, err)
		return
	}

	presenter.Response(w, &pb.UpdateResponse{
		Article: view.NewArticle(atl),
	})
}

func (c controller) delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		presenter.BadRequestError(w, err)
		return
	}

	atl := &model.Article{}
	if err := atl.Find(int64(id)); err != nil {
		c.log.Error(r.Context(), err)
		presenter.ApplicationException(w, err)
		return
	}
	if err := atl.Delete(); err != nil {
		c.log.Error(r.Context(), err)
		presenter.ApplicationException(w, err)
		return
	}

	atls := &model.Articles{}
	if err := atls.Find(); err != nil {
		c.log.Error(r.Context(), err)
		presenter.ApplicationException(w, err)
		return
	}

	presenter.Response(w, &pb.DeleteResponse{
		Articles: view.NewArticles(atls),
	})
}
