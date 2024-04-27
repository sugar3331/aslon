package core

import (
	"aslon/global"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Viper(path ...string) *viper.Viper {
	var config string

	config = "config/config.yaml"

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.GlobalConfig); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.GlobalConfig); err != nil {
		panic(err)
	}

	return v
}
