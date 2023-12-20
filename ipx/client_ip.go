package ipx

import (
	"context"
	"github.com/go-kratos/kratos/v2/metadata"
)

const (
	client_ip_key     = "x-forwarded-for"
	Default_client_ip = "1.1.1.1"
)

func GetClientIPFromContext(ctx context.Context) string {
	md, ok := metadata.FromClientContext(ctx)
	if !ok {
		return Default_client_ip
	}
	ip := md.Get(client_ip_key)
	if len(ip) == 0 {
		return Default_client_ip
	}
	return ip
}
