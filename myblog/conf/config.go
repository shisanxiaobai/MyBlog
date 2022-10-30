package conf

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var Config = new(AppConf)

type AppConf struct {
	ServerConfig `ini:"server"`
	DbConfig     `ini:"db"`
}
type ServerConfig struct {
	AppMode  string `ini:"AppMode"`
	HttpPort string `ini:"HttpPort"`
	JwtKey   string `ini:"JwtKey"`
	Username string `ini:"Username"`
	Password string `ini:"Password"`
}

type DbConfig struct {
	Driver     string `ini:"Driver"`
	Host       string `ini:"Host"`
	Port       string `ini:"Port"`
	User       string `ini:"User"`
	DbPassword string `ini:"DbPassword"`
	Name       string `ini:"Name"`
}

func init() {
	err := ini.MapTo(Config, "conf/config.ini")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
}
