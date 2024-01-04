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

	if ctx == nil {
		return Default_client_ip
	}

	md, ok := metadata.FromClientContext(ctx)
	if !ok || md == nil {
		return Default_client_ip
	}
	ip := md.Get(client_ip_key)
	if len(ip) == 0 {
		return Default_client_ip
	}
	return ip
}
