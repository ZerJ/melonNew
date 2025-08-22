package httpGet

import (
	"github.com/imroc/req/v3"
)

func HttpMelonQueryPost(client *req.Client, urls string, body map[string]string) (respByte []byte, err error) {
	resp, err := client.R().SetHeaders(map[string]string{
		"accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
		"accept-language": "zh-CN,zh;q=0.9",
		"cache-control":   "max-age=0",
		"content-type":    "application/x-www-form-urlencoded",
		"origin":          "https://tkglobal.melon.com",
		"priority":        "u=0, i",
		"referer":         "https://tkglobal.melon.com/reservation/popup/onestop.htm",

		"upgrade-insecure-requests": "1",
	}).SetFormData(body).Post(urls)
	return resp.ToBytes()
}
