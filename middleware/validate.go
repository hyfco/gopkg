package middleware

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

type Validate interface {
	ValidateAll() error
}

// Redacter defines how to log an object
type Redacter interface {
	Redact() string
}

// extractArgs returns the string of the req
func extractArgs(req interface{}) string {
	if redacter, ok := req.(Redacter); ok {
		return redacter.Redact()
	}
	if stringer, ok := req.(fmt.Stringer); ok {
		return stringer.String()
	}
	return fmt.Sprintf("%+v", req)
}

func ParameterValidate(logger log.Logger) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			var (
				kind      string
				operation string
			)

			if info, ok := transport.FromServerContext(ctx); ok {
				kind = info.Kind().String()
				operation = info.Operation()
			}

			if wrapper, ok := req.(Validate); ok {
				err = wrapper.ValidateAll()
				if err != nil {
					err = errors.New(400, "InvalidArgument", err.Error())
					return
				}
			} else {
				_ = log.WithContext(ctx, logger).Log(log.LevelInfo,
					"kind", "server",
					"component", kind,
					"operation", operation,
					"args", extractArgs(req),
				)
			}

			return handler(ctx, req)
		}
	}
}
