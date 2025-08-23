package ticket

import (
	"fmt"
	"github.com/imroc/req/v3"
	"melonNew/logging"
	"strings"
	"time"
)

func WSeqAuto(client *req.Client, aid string, interval time.Duration) (finalKey string, finalResult string, err error) {
	var ttl, key string
	var status, result string

	for {
		status, key, result, ttl, err = WSeqUnified(client, aid, key, ttl)
		if err != nil {
			return "", "", err
		}

		// 如果状态不是 201，说明请求成功或者失败，不再轮询
		if status != "201" {
			fmt.Println("------------1---------")
			fmt.Println(key)
			finalKey = key
			finalResult = result
			fmt.Println(finalKey)
			fmt.Println("------------2221---------")
			return finalKey, finalResult, nil
		}

		// 状态为 201，需要等待 ttl 或者固定间隔再轮询
		if interval > 0 {
			time.Sleep(interval)
		} else {
			// 默认 1 秒轮询
			time.Sleep(1 * time.Second)
		}
	}

	return finalKey, finalResult, nil
}

// WSeqUnified 发起 wseq 请求，初始或轮询校验都使用此方法
func WSeqUnified(client *req.Client, aid string, key string, ttl string) (status string, keyOut string, result string, ttlOut string, err error) {
	var url string
	if key == "" {
		// 初始请求
		gRtype := "5101"
		opcode := "5101"

		// 从 Cookie 中解析 keyCookie_T
		cookieStr := client.Cookies
		keyCookieT := ""
		for _, kv := range cookieStr {
			if kv.Name == "keyCookie_T" {
				keyCookieT = kv.Value
			}

		}
		if keyCookieT == "" {
			return "", "", "", "", fmt.Errorf("keyCookie_T not found in melonCookie")
		}

		url = fmt.Sprintf(
			"https://zam.melon.com/ts.wseq?opcode=%s&nfid=0&prefix=NetFunnel.gRtype=%s;&sid=service_1&aid=%s&js=yes&user_data=%s&%d",
			opcode, gRtype, aid, keyCookieT, time.Now().UnixNano()/1e6,
		)
	} else {
		// 轮询校验请求
		url = fmt.Sprintf(
			"https://zam.melon.com/ts.wseq?opcode=5002&key=%s&nfid=0&prefix=NetFunnel.gRtype=5002;&ttl=%s&sid=service_1&aid=%s&js=yes&%d",
			key, ttl, aid, time.Now().UnixNano()/1e6,
		)
	}

	// 统一请求头

	resp, err := client.R().
		SetHeader("Accept", "*/*").
		SetHeader("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6").
		SetHeader("Connection", "keep-alive").
		SetHeader("Referer", "https://tkglobal.melon.com/").
		SetHeader("Sec-Fetch-Dest", "script").
		SetHeader("Sec-Fetch-Mode", "no-cors").
		SetHeader("Sec-Fetch-Site", "same-site").
		Get(url)

	if err != nil {
		return "", "", "", "", err
	}

	body := resp.String()

	if !strings.Contains(body, "NetFunnel.gControl.result=") {
		return "", "", "", "", fmt.Errorf("unexpected response: %s", body)
	}
	fmt.Println(body)
	keyOut = strings.Split(strings.Split(body, ":key=")[1], "&nwait")[0]
	logging.Info(keyOut)
	logging.Info(body)
	result = strings.Split(strings.Split(body, "NetFunnel.gControl.result='")[1], "'; NetFunnel.gControl._showResult();")[0]

	if key == "" && strings.Contains(body, "gControl.result='5002:201") {
		ttlSplit := strings.Split(body, "&ttl=")
		if len(ttlSplit) > 1 {
			ttlOut = strings.Split(ttlSplit[1], "&ip")[0]
		}
		status = "201"
	} else {
		ttlSplit := strings.Split(body, "&ttl=")
		if len(ttlSplit) > 1 {
			ttlOut = strings.Split(ttlSplit[1], "&ip")[0]
		}
		status = strings.Split(strings.Split(body, "5002:")[1], ":key=")[0]
	}

	return status, keyOut, result, ttlOut, nil
}
