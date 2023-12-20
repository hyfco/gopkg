package dbx

const (
	defaultLimit = 100
)

func CalcPage[T int64 | int32 | int](page, pageSize T) (skip int, limit int) {
	if pageSize <= 0 || pageSize > defaultLimit {
		limit = defaultLimit
	}

	if page <= 0 {
		page = 1
	}

	skip = int(page-1) * limit
	return
}
