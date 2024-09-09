package personinn

import (
	"database/sql"
)

type PersonInn struct {
	ID       int    `json:"id"`
	PersonID int    `json:"person_id"`
	Name     string `json:"name"`
}

type PersonInnSql struct {
	ID       int            `json:"id"`
	PersonID sql.NullInt32  `json:"person_id"`
	Name     sql.NullString `json:"name"`
}

func (h *PersonInn) GetDBName() string {
	return `hr.person_inn`
}

func (h *PersonInn) GetID() int {
	return h.ID
}

func (h *PersonInn) SetID(id int) {
	h.ID = id
}

func (h *PersonInn) GetSQLInsert() string {
	return `INSERT INTO hr.person_inn (person_id, name) 
	VALUES ($1, $2) RETURNING (id);`
}

func (h *PersonInn) GetSQLGetEntity() string {
	return `SELECT id, person_id, name FROM hr.person_inn WHERE id = $1;`
}

func (h *PersonInn) GetSQLCountGetEntity() string {
	return `SELECT COUNT(*) FROM hr.person_inn WHERE id = $1;`
}

func (h *PersonInn) GetSQLUpdate() string {
	return `UPDATE hr.person_inn SET name = $1 WHERE id = $2 RETURNING (id);`
}
