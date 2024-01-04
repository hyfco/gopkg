package conf

import (
	"flag"
)

type FlagOptions struct {
	ConfigPath string
	Env        string
}

func NewFlagOptions() *FlagOptions {
	ret := &FlagOptions{}
	flag.StringVar(&ret.ConfigPath, "conf", "../config", "config path. e.g: --conf ./config/config_dev.yaml")
	flag.StringVar(&ret.Env, "env", "dev", "application run env")
	flag.Parse()

	return ret
}
