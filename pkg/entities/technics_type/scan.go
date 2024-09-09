package personinn

import "github.com/jackc/pgx/v5"

func (lcs *PersonInn) ScanEntity(row pgx.Rows) error {
	var lcsSQL PersonInnSql

	if err := row.Scan(
		&lcsSQL.ID,
		&lcsSQL.PersonID,
		&lcsSQL.Name,
	); err != nil {
		return err
	}

	lcs.ID = lcsSQL.ID

	if lcsSQL.PersonID.Valid {
		lcs.PersonID = int(lcsSQL.PersonID.Int32)
	}

	if lcsSQL.Name.Valid {
		lcs.Name = lcsSQL.Name.String
	}

	return nil
}

func (h *PersonInnList) ScanEntityCollection(rows pgx.Rows) error {
	i := 0
	x := *&h.List

	for rows.Next() {
		if err := x[i].ScanEntity(rows); err != nil {
			return err
		}
		i++
	}
	return nil
}
