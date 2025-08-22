package main

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"melonNew/global"
	"melonNew/login"
	"melonNew/member"
	"melonNew/ticket"
	"os"
)

func main() {
	configPath := flag.String("config", "", "config file path or directory")
	flag.Parse()

	global.GlobalConfig = initConfig(*configPath)
	go dynamicConfig()
	email := global.GlobalConfig.GetString("email")

	pwd := global.GlobalConfig.GetString("pwd")
	proxy := global.GlobalConfig.GetString("proxy")
	res, err := login.LoginMelonWithReq(email, pwd, proxy)
	if err != nil {
		log.Fatal(err)
	}
	member.GetMember(res.Client)
	ticket.GetTicket(res.Client)
}
func initConfig(configPath string) *viper.Viper {
	cfg := viper.New()

	if configPath == "" {
		configPath = "./config.yaml" // 默认路径
	}

	info, err := os.Stat(configPath)
	if err != nil {
		fmt.Println("配置文件路径无效:", configPath)
		return cfg
	}

	if info.IsDir() {
		cfg.SetConfigName("config")
		cfg.AddConfigPath(configPath)
		cfg.SetConfigType("yaml")
	} else {
		cfg.SetConfigFile(configPath)
	}

	err = cfg.ReadInConfig()
	if err != nil {
		fmt.Println("读取配置失败:", err)
		return cfg
	}

	fmt.Println("已加载配置文件:", cfg.ConfigFileUsed())
	return cfg
}

func dynamicConfig() {
	global.GlobalConfig.WatchConfig()
	global.GlobalConfig.OnConfigChange(func(event fsnotify.Event) {
		fmt.Printf("Detect config change: %s \n", event.String())
	})
}
