package goenv

import (
	"github.com/spf13/viper"
)

var Env *viper.Viper

// Environment Initialization
func Init(env string, parser string, path string) error {
	var err error
	Env = viper.New()
	Env.SetConfigType(parser)
	Env.SetConfigName(env)
	Env.AddConfigPath(path + "/")
	err = Env.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}

// Load config
func LoadEnv() *viper.Viper {
	return Env
}
