package stat

type Config struct {
	Addr string `mapstructure:"addr" json:"addr" yaml:"addr" toml:"addr"`

	Metrics    metrics    `mapstructure:"metrics"  json:"metrics"    yaml:"metrics"     toml:"metrics"`
	Profiler   profiler   `mapstructure:"profiler" json:"profiler"   yaml:"profiler"    toml:"profiler"`
	Probe      probe      `mapstructure:"probe"    json:"probe"      yaml:"probe"       toml:"probe"`
	Prometheus prometheus `mapstructure:"stat"     json:"prometheus" yaml:"prometheus"  toml:"prometheus"`
}

type metrics struct {
	SwitchOff bool `mapstructure:"switch_off" json:"switch_off" yaml:"switch_off" toml:"switch_off"`
}

type profiler struct {
	SwitchOff bool `mapstructure:"switch_off" json:"switch_off" yaml:"switch_off" toml:"switch_off"`
}

type probe struct {
	SwitchOff bool `mapstructure:"switch_off" json:"switch_off" yaml:"switch_off" toml:"switch_off"`
}

type prometheus struct {
	SwitchOff bool `mapstructure:"switch_off" json:"switch_off" yaml:"switch_off" toml:"switch_off"`
}
