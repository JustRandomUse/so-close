package personinn

import (
	"context"
	"encoding/json"

	"dev.kodeks.net/Party/back_frame/common/postgres"
	"go.uber.org/zap"
)

type PersonInnController struct {
	PgDB *postgres.PgDB
}

func NewController(db *postgres.PgDB) *PersonInnController {

	return &PersonInnController{
		PgDB: db,
	}
}

func (ctrl *PersonInnController) GetPersonInn(ctx context.Context, id int) (*PersonInn, error) {

	const fname = `GetPersonInn`

	LabConSt := &PersonInn{
		ID: id,
	}

	err := ctrl.PgDB.GetEntity(ctx, LabConSt)
	if err != nil {
		zap.S().Debugf("%s  GetEntity error: %v", fname, err)
		return nil, err
	}

	return LabConSt, err
}

func (ctrl *PersonInnController) CreatePersonInn(ctx context.Context, Entity *PersonInn) (int, error) {
	return ctrl.PgDB.CreateEntity(ctx, Entity, Entity.PersonID, Entity.Name)
}

func (ctrl *PersonInnController) UpdatePersonInn(ctx context.Context, Entity *PersonInn) (int, error) {
	return ctrl.PgDB.UpdateEntity(ctx, Entity, Entity.Name, Entity.ID)
}

func (ctrl *PersonInnController) listPersonInn(ctx context.Context, filter *PersonInnFilter) (*PersonInnList, error) {

	const fname = `listPersonInn`

	var err error
	var tmp []byte

	if tmp, err = json.MarshalIndent(filter, "", " "); err != nil {
		zap.S().Debug("postgres " + fname + " Error filter: " + string(tmp))
	}

	list := &PersonInnList{
		TotalCount: 0,
		Len:        0,
		Offset:     filter.Offset,
		List:       nil,
		Filter:     *filter,
	}
	if err = ctrl.PgDB.ListEntity(ctx, list); err != nil {
		zap.S().Debug(fname + "ListEntity Error: " + err.Error())
	}

	return list, err
}

func (ctrl *PersonInnController) DeletePersonInn(ctx context.Context, id int) error {
	const dbname = `hr.person_inn`
	return ctrl.PgDB.RealDeleteObject(ctx, id, dbname)
}
