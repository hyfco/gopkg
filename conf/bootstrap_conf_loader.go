package conf

import (
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
)

func NewBootstrapConfigFromPath(filePath string) *Bootstrap {
	loader := config.New(config.WithSource(file.NewSource(filePath)))
	if err := loader.Load(); err != nil {
		panic(err)
	}
	var (
		conf = new(Bootstrap)
	)

	if err := loader.Scan(conf); err != nil {
		panic(err)
	}
	return conf
}
