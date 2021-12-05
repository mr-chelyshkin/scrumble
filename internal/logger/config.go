package logger

type Config struct {
	Level   string `mapstructure:"level" json:"level" yaml:"level" toml:"level"`
	LogPath string `mapstructure:"path"  json:"path"  yaml:"path"  toml:"path"`

	MaxSize    int `mapstructure:"log_size"    json:"log_size"    yaml:"log_size"    toml:"log_size"`
	MaxAge     int `mapstructure:"log_age"     json:"log_age"     yaml:"log_age"     toml:"log_age"`
	MaxBackups int `mapstructure:"log_backups" json:"log_backups" yaml:"log_backups" toml:"log_backups"`
}
