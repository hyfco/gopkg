package middleware

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
)

type Validate interface {
	ValidateAll() error
}

func ParameterValidate() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if wrapper, ok := req.(Validate); ok {
				err = wrapper.ValidateAll()
				if err != nil {
					err = errors.New(400, "InvalidArgument", err.Error())
					return
				}
			}
			return handler(ctx, req)
		}
	}
}
