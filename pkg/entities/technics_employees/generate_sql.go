package personinn

import (
	"strconv"

	"dev.kodeks.net/Party/back_frame/common/postgres"
	"go.uber.org/zap"
)

func generateSQLListPersonInn(filter *PersonInnFilter) postgres.SqlParts {
	const fname = "generateSQLListPersonInn"
	const And = ` AND `
	const dbname = ` hr.person_inn `
	zap.S().Debug(fname)
	var res postgres.SqlParts

	res.Head = ` SELECT `
	res.Count = ` COUNT (*) `
	res.Fields = ` id, person_id, name `
	res.Body = ` FROM ` + dbname

	/*
		Limit             int    `json:"limit"`
		Offset            int    `json:"offset"`
	*/
	if filter.Limit != 0 {
		res.Tail += ` LIMIT ` + strconv.Itoa(filter.Limit)
	}

	if filter.Offset != 0 {
		res.Tail += ` OFFSET ` + strconv.Itoa(filter.Offset)
	}

	return res
}
