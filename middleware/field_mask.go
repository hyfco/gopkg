package middleware

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/hyfco/gopkg/fieldmask"
	"google.golang.org/protobuf/proto"
)

func ResponseFieldMask(visibility fieldmask.Visibility) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			reply, err = handler(ctx, req)
			if err != nil {
				return
			}
			if resp, ok := reply.(proto.Message); ok {
				fieldmask.Filter(resp, visibility)
			}
			return
		}
	}
}
