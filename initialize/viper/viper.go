package viper

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/kingdowliu/go-web-template/global"
	"github.com/spf13/viper"
)

func InitViper(path ...string) {
	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "f", "config.yaml", "配置文件地址")
		flag.Parse()
		if len(config) == 0 {
			panic("配置文件缺失")
		}
	} else {
		config = path[0]
	}

	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigFile(config)

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("Fatal error config file: %s \n", err))
	}

	// 监听配置变更
	v.WatchConfig()

	// 配置改变，重新解析
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		if err := v.Unmarshal(&global.DB_CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	if err := v.Unmarshal(&global.DB_CONFIG); err != nil {
		fmt.Println(err)
	}

	global.VIPER = v
}
