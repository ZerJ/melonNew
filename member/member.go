package member

import (
	"encoding/json"
	"fmt"
	"github.com/imroc/req/v3"
	"io/ioutil"
	"melonNew/global"
	"melonNew/model"
)

func GetMember(client *req.Client) (err error) {
	u := "https://gapi.melon.com/v1/gmember/member"
	resp, err := client.ImpersonateChrome().R().
		// 不用手动覆盖 UA / sec-ch-ua，用 ImpersonateChrome 默认值即可
		SetHeaders(map[string]string{
			"accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
			"accept-language": "zh-CN,zh;q=0.9",
			"cache-control":   "max-age=0",
			"content-type":    "application/json",
			"origin":          "https://m.globalticket.melon.com",
			"priority":        "u=0, i",
			"referer":         "https://m.globalticket.melon.com/",
		}).
		Get(u)
	fmt.Println(resp.StatusCode)

	respByte, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return
	}
	var r model.GetMember
	err = json.Unmarshal(respByte, &r)
	if err != nil {
		return
	}
	if r.Status == "0000" {
		global.GlobalConfig.Set("userName", r.Data.Name)
		global.GlobalConfig.Set("email", r.Data.Email)
		fmt.Println(r.Data)
	}
	return
}
