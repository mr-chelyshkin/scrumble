package torrent_fetcher

type Config struct {
	WatchDir    string `mapstructure:"watch_dir"    json:"watch_dir"    yaml:"watch_dir"    toml:"watch_dir"`
	CompleteDir string `mapstructure:"complete_dir" json:"complete_dir" yaml:"complete_dir" toml:"complete_dir"`
	DownloadDir string `mapstructure:"download_dir" json:"download_dir" yaml:"download_dir" toml:"download_dir"`
}
