package core

import (
	"tsetmc-gopher/global"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Viper(configfilePath string) *viper.Viper {
	v := viper.New()
	v.SetConfigFile(configfilePath)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.BRC_CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	if err := v.Unmarshal(&global.BRC_CONFIG); err != nil {
		fmt.Println(err)
	}
	return v
}
