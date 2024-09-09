package personinn

import (
	"dev.kodeks.net/Party/back_frame/common/config"
	"dev.kodeks.net/Party/back_frame/common/utils"
)

func normalizePersonInnFilter(filter *PersonInnFilter) {
	filter.Limit = utils.Abs(filter.Limit)
	filter.Offset = utils.Abs(filter.Offset)

	if filter.Limit > config.MAX_FILTER_LIMIT || filter.Limit == 0 {
		filter.Limit = config.MAX_FILTER_LIMIT
	}
}
