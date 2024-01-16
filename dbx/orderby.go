package dbx

import (
	"strings"

	"gorm.io/gen/field"
)

type GetFieldByName interface {
	GetFieldByName(fieldName string) (field.OrderExpr, bool)
}

func ConvertOrderBy(table GetFieldByName, orderBy string) []field.Expr {
	if table == nil || orderBy == "" {
		return []field.Expr{}
	}
	orders := strings.Split(orderBy, ",")
	fields := make([]string, len(orders))
	if len(orders) == 0 {
		return nil
	}

	ret := []field.Expr{}

	for _, f := range fields {
		if strings.Contains(f, "+") {
			f = strings.ReplaceAll(f, "+", "")
			e, ok := table.GetFieldByName(f)
			if ok {
				ret = append(ret, e.Asc())
			}
		} else if strings.Contains(f, "-") {
			f = strings.ReplaceAll(f, "-", "")
			e, ok := table.GetFieldByName(f)
			if ok {
				ret = append(ret, e.Desc())
			}
		} else {
			e, ok := table.GetFieldByName(f)
			if ok {
				ret = append(ret, e.Asc())
			}
		}
	}

	return ret
}
