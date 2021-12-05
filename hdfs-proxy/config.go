package hdfs_proxy

import "fmt"

type Config struct {
	AclPath string `mapstructure:"ppp" json:"ppp" yaml:"ppp"`
	//Config config.Config
}

func (c Config) validate(cfg interface{}) error {
	fmt.Println(cfg)
	return nil
}