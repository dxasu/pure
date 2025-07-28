package storage

import (
	"os"
	"testing"

	"github.com/spf13/viper"
)

func TestData_GetString(t *testing.T) {
	type fields struct {
		v *viper.Viper
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Data{
				v: tt.fields.v,
			}
			v := tt.fields.v
			v.SetDefault("mysql.port", 3306)   // 默认值
			os.Setenv("DB_HOST", "10.0.0.1")   // 环境变量
			v.AutomaticEnv()                   // 自动加载环境变量
			v.BindEnv("mysql.host", "DB_HOST") // 显式绑定
			if got := d.GetString(tt.args.key); got != tt.want {
				t.Errorf("Data.GetString() = %v, want %v", got, tt.want)
			}
		})
	}
}
