package personinn

import "dev.kodeks.net/Party/back_frame/common/postgres"

type PersonInnList struct {
	TotalCount int             `json:"total_count"`
	Offset     int             `json:"offset"`
	Len        int             `json:"len"`
	List       []PersonInn     `json:"person_inn_list"`
	Filter     PersonInnFilter `json:"-"`
}

type PersonInnFilter struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

func (hl *PersonInnList) MakeListEntity(c int) {
	hl.List = make([]PersonInn, c)
	hl.Len = c
}

func (hl *PersonInnList) GetSql() postgres.SqlParts {
	return generateSQLListPersonInn(&hl.Filter)
}

func (hl *PersonInnList) GetFilterOffset() int {
	return hl.Filter.Offset
}

func (hl *PersonInnList) GetFilterLimit() int {
	return hl.Filter.Limit
}

func (hl *PersonInnList) SetFilterOffset(o int) {
	hl.Len = 0
}

func (hl *PersonInnList) SetTotalCount(c int) {
	hl.TotalCount = c
}

func (hl *PersonInnList) SetOffset(c int) {
	hl.Offset = c
}

func (hl *PersonInnList) SetListLen(o int) {
	hl.Len = o
}
