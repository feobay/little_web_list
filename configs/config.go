package configs

import (
	"github.com/spf13/viper"
	"sync"
)

var (
	globalConfig GlobalConf
	once         sync.Once
)

type GlobalConf struct {
	DbConfig DbConf `yaml:"db" mapstructure:"db"`
}

type DbConf struct {
	Host        string `yaml:"host" mapstructure:"host"`
	Port        int    `yaml:"port" mapstructure:"port"`
	User        string `yaml:"user" mapstructure:"user"`
	PassWord    string `yaml:"password" mapstructure:"password"`
	Dbname      string `yaml:"dbname" mapstructure:"dbname"`
	MaxIdleConn int    `yaml:"max_idle_conn" mapstructure:"max_idle_conn"`
	MaxOpenConn int    `yaml:"max_open_conn" mapstructure:"max_open_conn"`
	MaxIdleTime int    `yaml:"max_idle_time" mapstructure:"max_idle_time"`
}

func GetGlobalConfig() *GlobalConf {
	once.Do(readConf)
	return &globalConfig
}

func readConf() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./configs")

	err := viper.ReadInConfig()
	if err != nil {
		panic("viper.ReadInConfig Failed.")
	}

	err = viper.Unmarshal(&globalConfig)
	if err != nil {
		panic("viper.Unmarshal Failed")
	}
}
