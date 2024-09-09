package personinn

import (
	"encoding/json"
	"net/http"
	"strconv"

	"dev.kodeks.net/Party/back_frame/common/postgres"
	"dev.kodeks.net/Party/back_frame/common/utils"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type PersonInnRouter struct {
	ctrlr *PersonInnController
}

func New(db *postgres.PgDB) *PersonInnRouter {
	hdb := NewController(db)

	return &PersonInnRouter{
		ctrlr: hdb,
	}
}

func (wr *PersonInnRouter) PersonInnRouter() chi.Router {
	r := chi.NewRouter()
	r.Get("/", wr.listPersonInn)
	r.Post("/list", wr.listPersonInn)
	r.Post("/", wr.createPersonInn)

	r.Route("/{Id}", func(r chi.Router) {
		r.Get("/", wr.getPersonInn)
		r.Put("/", wr.updatePersonInn)
		r.Delete("/", wr.deletePersonInn)

	})

	return r
}

func (wr *PersonInnRouter) listPersonInn(w http.ResponseWriter, r *http.Request) {

	const fname = "listPersonInn"

	ctx := r.Context()
	filter := &PersonInnFilter{}

	var err error

	if err = json.NewDecoder(r.Body).Decode(filter); err != nil {
		zap.S().Debugf("%s %v Error: %v", fname, filter, err)
	}

	normalizePersonInnFilter(filter)

	mngrs, err := wr.ctrlr.listPersonInn(ctx, filter)
	if err != nil {
		zap.S().Debugf("%s error: %v", fname, err)
		utils.RenderJSON(w, fname+" error: "+err.Error(), http.StatusBadRequest)
		return
	}

	utils.RenderJSON(w, mngrs, http.StatusOK)
}

func (wr *PersonInnRouter) getPersonInn(w http.ResponseWriter, r *http.Request) {

	const fname = "getPersonInn"

	ctx := r.Context()
	idStr := chi.URLParam(r, "Id")

	var err error
	var id int

	if id, err = strconv.Atoi(idStr); err != nil {
		zap.S().Debugf("%s id = %s, Error: %v", fname, idStr, err)
		utils.RenderJSON(w, fname+" id = "+idStr+" Error: "+err.Error(), http.StatusBadRequest)
		return
	}

	mngr, err := wr.ctrlr.GetPersonInn(ctx, id)
	if err != nil {
		zap.S().Debugf("Repo %s error: %v", fname, err)
		utils.RenderJSON(w, "Repo "+fname+" error: "+err.Error(), http.StatusBadRequest)
		return
	}

	utils.RenderJSON(w, mngr, http.StatusOK)
}

func (wr *PersonInnRouter) createPersonInn(w http.ResponseWriter, r *http.Request) {

	const fname = `createPersonInn`

	ctx := r.Context()
	mngr := &PersonInn{}

	var err error

	if err = json.NewDecoder(r.Body).Decode(mngr); err != nil {
		zap.S().Debugf(fname+": %v Error: %v", mngr, err)
		utils.RenderJSON(w, fname+" error: "+err.Error(), http.StatusBadRequest)
		return
	}

	if mngr.ID, err = wr.ctrlr.CreatePersonInn(ctx, mngr); err != nil {
		zap.S().Debugf("%s: %v\nError: %v", fname, mngr, err)
		utils.RenderJSON(w, fname+" error: "+err.Error(), http.StatusBadRequest)
		return
	}

	utils.RenderJSON(w, mngr.ID, http.StatusOK)
}

func (wr *PersonInnRouter) updatePersonInn(w http.ResponseWriter, r *http.Request) {

	const fname = "updatePersonInn"

	ctx := r.Context()
	idStr := chi.URLParam(r, "Id")

	var err error
	var id int

	if id, err = strconv.Atoi(idStr); err != nil {
		zap.S().Debugf(fname+" id = %s, Error: %v", idStr, err)
		utils.RenderJSON(w, fname+" id = "+idStr+" Error: "+err.Error(), http.StatusBadRequest)
		return
	}

	prsn := &PersonInn{}

	if err = json.NewDecoder(r.Body).Decode(prsn); err != nil {
		zap.S().Debugf(fname+": %v Error: %v", prsn, err)
		utils.RenderJSON(w, fname+" error: "+err.Error(), http.StatusBadRequest)
		return
	}

	prsn.ID = id

	if prsn.ID, err = wr.ctrlr.UpdatePersonInn(ctx, prsn); err != nil {
		zap.S().Debugf("%s: %v\nError: %v", fname, prsn, err)
		utils.RenderJSON(w, fname+" error: "+err.Error(), http.StatusBadRequest)
		return
	}

	utils.RenderJSON(w, prsn, http.StatusOK)
}

func (wr *PersonInnRouter) deletePersonInn(w http.ResponseWriter, r *http.Request) {

	const fname = `deletePersonInn`

	ctx := r.Context()
	idStr := chi.URLParam(r, "Id")

	var err error
	var id int

	if id, err = strconv.Atoi(idStr); err != nil {
		zap.S().Debugf("%s id = %s, Error: %v", fname, idStr, err)
		utils.RenderJSON(w, fname+" id = "+idStr+" Error: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err = wr.ctrlr.DeletePersonInn(ctx, id); err != nil {
		zap.S().Debugf("%s error: %v", fname, err)
		utils.RenderJSON(w, fname+" error: "+err.Error(), http.StatusBadRequest)
		return
	}

	utils.RenderJSON(w, "", http.StatusOK)
}
