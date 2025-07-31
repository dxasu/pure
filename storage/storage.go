package storage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var ErrEmptyViper = fmt.Errorf("viper instance is nil")

type Data struct {
	fp string
	v  *viper.Viper
}

func (d *Data) GetString(key string) string {
	if d.v == nil {
		panic(ErrEmptyViper)
	}
	value := d.v.GetString(key)
	return value
}

func (d *Data) GetInt(key string) int {
	if d.v == nil {
		panic(ErrEmptyViper)
	}
	value := d.v.GetInt(key)
	return value
}

func (d *Data) Get(key string) (any, error) {
	if d.v == nil {
		panic(ErrEmptyViper)
	}
	value := d.v.Get(key)
	if value == nil {
		return nil, fmt.Errorf("key %s not found", key)
	}
	return value, nil
}

func (d *Data) Set(key string, value any) error {
	if d.v == nil {
		panic(ErrEmptyViper)
	}
	d.v.Set(key, value)
	return nil
}

func (d *Data) Delete(key string) error {
	if d.v == nil {
		panic(ErrEmptyViper)
	}
	if !d.v.IsSet(key) {
		return fmt.Errorf("key %s not found", key)
	}
	d.v.Set(key, nil)
	return nil
}

func (d *Data) Save() error {
	if d.v == nil {
		panic(ErrEmptyViper)
	}

	if err := d.v.WriteConfig(); err != nil {
		return fmt.Errorf("failed to write config err: %w", err)
	}
	return nil
}

func (d *Data) GetFilePath() string {
	if d.v == nil {
		panic(ErrEmptyViper)
	}
	return d.fp
}

func (d *Data) GetViper() *viper.Viper {
	if d.v == nil {
		panic(ErrEmptyViper)
	}
	return d.v
}

func (d *Data) PrintJSON() {
	j := d.v.AllSettings()
	data, _ := json.MarshalIndent(j, "", "  ")
	fmt.Println("\033[32mJson:\033[0m")
	fmt.Println(string(data))
}

func (d *Data) JSON() string {
	j := d.v.AllSettings()
	data, _ := json.Marshal(j)
	return string(data)
}

func (d *Data) Init(name string, dPath *DataPath) error {
	filePath := dPath.GetFilePath(name)
	d.fp = filePath
	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			// If the file does not exist, we can create a new viper instance
			d.v = viper.New()
			d.v.SetConfigFile(filePath)
			d.v.AddConfigPath(dPath.GetDir())
			return nil
		}
		return fmt.Errorf("failed to read data from %s: %w", filePath, err)
	}
	if len(data) == 0 {
		return fmt.Errorf("file %s is empty", filePath)
	}
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}

	v := viper.New()
	v.SetConfigFile(filePath)
	v.SetConfigType("json")

	err = v.ReadConfig(bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("failed to parse JSON from %s: %w", filePath, err)
	}
	d.v = v
	return nil
}
