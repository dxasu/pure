package storage

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Data struct {
	v *viper.Viper
}

func (d *Data) GetString(key string) (string, error) {
	if d.v == nil {
		return "", fmt.Errorf("viper instance is nil")
	}
	value := d.v.GetString(key)
	if value == "" && !d.v.IsSet(key) {
		return "", fmt.Errorf("key %s not found", key)
	}
	return value, nil
}

func (d *Data) GetInt(key string) (int, error) {
	if d.v == nil {
		return 0, fmt.Errorf("viper instance is nil")
	}
	value := d.v.GetInt(key)
	if value == 0 && !d.v.IsSet(key) {
		return 0, fmt.Errorf("key %s not found", key)
	}
	return value, nil
}

func (d *Data) Get(key string) (any, error) {
	if d.v == nil {
		return nil, fmt.Errorf("viper instance is nil")
	}
	value := d.v.Get(key)
	if value == nil {
		return nil, fmt.Errorf("key %s not found", key)
	}
	return value, nil
}

func (d *Data) Set(key string, value any) error {
	if d.v == nil {
		return fmt.Errorf("viper instance is nil")
	}
	d.v.Set(key, value)
	return nil
}

func (d *Data) Delete(key string) error {
	if d.v == nil {
		return fmt.Errorf("viper instance is nil")
	}
	if !d.v.IsSet(key) {
		return fmt.Errorf("key %s not found", key)
	}
	d.v.Set(key, nil)
	return nil
}

func (d *Data) Save() error {
	if d.v == nil {
		return fmt.Errorf("viper instance is nil")
	}

	if err := d.v.WriteConfig(); err != nil {
		return fmt.Errorf("failed to write config err: %w", err)
	}
	return nil
}

func NewData(name string) (*Data, error) {
	tempDir := os.TempDir()
	filePath := filepath.Join(tempDir, name)
	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			// If the file does not exist, we can create a new viper instance
			v := viper.New()
			v.SetConfigName(filePath)
			v.AddConfigPath(tempDir)
			return &Data{v: v}, nil
		}
		return nil, fmt.Errorf("failed to read data from %s: %w", filePath, err)
	}
	if len(data) == 0 {
		return nil, fmt.Errorf("file %s is empty", filePath)
	}
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}

	v := viper.New()
	v.SetConfigType("json")
	err = v.ReadConfig(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON from %s: %w", filePath, err)
	}

	return &Data{v: v}, nil
}
