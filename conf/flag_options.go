package conf

import "flag"

type FlagOptions struct {
	ConfigPath string
	Env        string
}

func NewFlagOptions() *FlagOptions {
	ret := &FlagOptions{}
	flag.StringVar(&ret.ConfigPath, "conf", "../configs", "config path. e.g: -conf ../config.yaml")
	flag.StringVar(&ret.Env, "env", "dev", "application run env")
	return ret
}
