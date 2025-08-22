package main

import (
	"fmt"
	"log"
	"melonNew/global"
	"melonNew/login"
)

func main() {

	email := global.GlobalConfig.GetString("email")

	pwd := global.GlobalConfig.GetString("pwd")
	proxy := global.GlobalConfig.GetString("proxy")
	res, err := login.LoginMelonWithReq(email, pwd, proxy)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
