package config

import (
	"github.com/spf13/viper"
)

// Viper 配置实例
var v *viper.Viper

// New 构造函数
func New() *viper.Viper {
	return v
}

// Init 初始化配置
func Init(name string) {
	v = &viper.Viper{}

	v.AutomaticEnv()
	v.SetEnvPrefix(name)
}
