package hdfs_proxy

import (
	"fmt"
	"github.com/pkg/errors"
)

type Config struct {
	AclPath string `mapstructure:"ppp" json:"ppp" yaml:"ppp"`
	//Config config.Config
}

func (c Config) validate(cfg interface{}) error {
	fmt.Println(cfg)
	return errors.Errorf("Qwerty ERRRORRS")
}