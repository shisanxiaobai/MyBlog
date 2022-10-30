package conf

import (
	"fmt"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// viper实时监听配置
func init() {
	v := viper.New()
	dir, _ := os.Getwd()
	v.AddConfigPath(dir + "/conf")
	v.SetConfigName("config")
	v.SetConfigType("ini")

	err := v.ReadInConfig()
	if err != nil {
		log.Print("读取配置失败")
		panic(err)
	}

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Println("Config file changed:", e.Name)
	})

}
