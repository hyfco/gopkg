package mapstructurehookfunc

import (
	"github.com/mitchellh/mapstructure"
	"reflect"
	"time"
)

func TimestampToTimeHookFunc() mapstructure.DecodeHookFunc {
	return func(f, t reflect.Type, data interface{}) (interface{}, error) {
		if f.Kind() == reflect.Int64 && t == reflect.TypeOf(time.Time{}) {
			return time.Unix(data.(int64), 0), nil
		}

		return data, nil
	}
}
