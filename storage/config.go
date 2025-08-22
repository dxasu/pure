package storage

import (
	"os"
	"path/filepath"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

type config struct {
	viper *viper.Viper
}

func New(configPath string, watch bool) (*config, error) {
	v := viper.New()
	v.SetConfigFile(configPath)
	v.SetConfigType("json")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		dir := filepath.Dir(configPath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return nil, err
		}
		if err := v.WriteConfig(); err != nil {
			return nil, err
		}
	} else {
		if err := v.ReadInConfig(); err != nil {
			return nil, err
		}
	}

	if watch {
		v.WatchConfig()
	}
	return &config{viper: v}, nil
}

func (c *config) Range(key string, fn func(k string, v any)) {
	m, err := cast.ToStringMapE(c.viper.Get(key))
	if err == nil {
		for k, v := range m {
			fn(k, v)
		}
	} else {
		s := cast.ToSlice(c.viper.Get(key))
		for k, v := range s {
			fn(cast.ToString(k), v)
		}
	}
}

func (c *config) Set(kv ...string) {
	for i := 0; i < len(kv)>>1<<1; i += 2 {
		c.viper.Set(kv[i], kv[i+1])
	}
	c.viper.WriteConfig()
}

func (c *config) Get(k string) any {
	return c.viper.Get(k)
}
