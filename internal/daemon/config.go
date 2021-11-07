package daemon

type Config struct {
	NoStat  bool `mapstructure:"no_stat" json:"no_stat" yaml:"no_stat"`
}
