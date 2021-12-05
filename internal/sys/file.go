package sys

import (
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"sync"
	"syscall"

	"github.com/fsnotify/fsnotify"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func CurrentUsername() string {
	current, err := user.Current()
	if err != nil {
		return "unknown"
	}
	return current.Username
}

func IsWritable(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	if !info.IsDir() {
		return false, errors.Errorf("%s is not directory", path)
	}
	if info.Mode().Perm()&(1<<(uint(7))) == 0 {
		return false, errors.Errorf("%s write permission bit is not set for user: %s", path, CurrentUsername())
	}

	var stat syscall.Stat_t
	if err = syscall.Stat(path, &stat); err != nil {
		return false, errors.Errorf("unable to get stat for %s", path)
	}
	if uint32(os.Geteuid()) != stat.Uid {
		return false, errors.Errorf("user '%s' doesn't have permission to write to %s", CurrentUsername(), path)
	}

	return true, err
}

// TODO: cfg interface -> interface with method validate()
func ParseFile(path string, cfg interface{}) error {
	dir := filepath.Dir(path)
	ext := strings.TrimPrefix(filepath.Ext(path), ".")
	file := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))

	viper.SetConfigName(file)
	viper.AddConfigPath(dir)
	viper.SetConfigType(ext)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(&cfg, func(config *mapstructure.DecoderConfig) {
		config.TagName = ext
	}); err != nil {
		return err
	}

	return nil
}

func ParseFileOnChange(path string, cfg interface{}, validate func(cfg interface{}) error) error {
	var read = func(
		cfg interface{},
		v *viper.Viper,
		ext string,
		mu *sync.Mutex,
		errCh chan error,
		validate func(cfg interface{}) error,
	) {

		if err := v.ReadInConfig(); err != nil {
			errCh <- err
			return
		}

		var data interface{}
		if err := v.Unmarshal(&data, func(config *mapstructure.DecoderConfig) {
			config.TagName = ext
		}); err != nil {
			errCh <- err
			return
		}
		if err := validate(data); err != nil {
			errCh <- err
			return
		}

		mu.Lock()
		cfg = data
		mu.Unlock()
	}

	dir := filepath.Dir(path)
	ext := strings.TrimPrefix(filepath.Ext(path), ".")
	file := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))

	errCh := make(chan error, 1)
	mu := &sync.Mutex{}

	v := viper.New()
	v.SetConfigName(file)
	v.AddConfigPath(dir)
	v.SetConfigType(ext)
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		read(cfg, v, ext, mu, errCh, validate)
	})

	return <-errCh
}
