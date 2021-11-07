package stat

type Config struct {
	Addr string `mapstructure:"addr" json:"addr" yaml:"addr"`

	Metrics    metrics    `mapstructure:"metrics"  json:"metrics"    yaml:"metrics"`
	Profiler   profiler   `mapstructure:"profiler" json:"profiler"   yaml:"profiler"`
	Probe      probe      `mapstructure:"probe"    json:"probe"      yaml:"probe"`
	Prometheus prometheus `mapstructure:"stat"     json:"prometheus" yaml:"prometheus"`
}

type metrics struct {
	SwitchOff bool `mapstructure:"switch_off" json:"switch_off" yaml:"switch_off"`
}

type profiler struct {
	SwitchOff bool `mapstructure:"switch_off" json:"switch_off" yaml:"switch_off"`
}

type probe struct {
	SwitchOff bool `mapstructure:"switch_off" json:"switch_off" yaml:"switch_off"`
}

type prometheus struct {
	SwitchOff bool `mapstructure:"switch_off" json:"switch_off" yaml:"switch_off"`
}
