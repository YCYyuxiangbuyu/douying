package core

import (
	"douyin/common/config"
	"douyin/common/global"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const filename = "./settings.yaml"

func InitConf() {
	var c = &config.Config{}

	viper.SetConfigFile(filename) //viper是Go应用程序的完整配置解决方案支持从JSON，TOML，YAML，HCL和Java等属性配置文件中读取,这里是yaml
	//
	if err := viper.ReadInConfig(); err != nil { //读取文件
		fmt.Println("viper.ReadInConfig Failed, err : ", err.Error())
		return
	}
	if err := viper.Unmarshal(c); err != nil { //反序列化 你可以选择将所有或特定的值解析到结构体、map等,解析到c中
		fmt.Println("viper.Unmarshal Failed, err : ", err.Error())
		return //Viper在后台使用github.com/mitchellh/mapstructure来解析值，其默认情况下使用mapstructuretag。

		//注意 当我们需要将viper读取的配置反序列到我们定义的结构体变量中时，一定要使用mapstructuretag哦！
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件配修改了")
		if err := viper.Unmarshal(c); err != nil { //配置文件发生变化后要同步到全局变量global.Config = c
			fmt.Println("viper.Unmarshal Failed, err : ", err.Error())
			return
		}

	}) //实现配置的热更新，不用重启项目新配置即可生效

	global.Config = c
	return
}
