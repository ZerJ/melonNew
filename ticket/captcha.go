package ticket

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/anaskhan96/soup"
	"github.com/imroc/req/v3"
	"io/ioutil"
	"melonNew/logging"
	"melonNew/model"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func prodKey(client *req.Client, prodId, scheduleNos string) (r model.ProdKey) {
	fmt.Println("prodKey")
	u := "https://tkglobal.melon.com/tktapi/glb/product/prodKey.json?prodId=" + prodId + "&scheduleNo=" + scheduleNos + "&v=1"
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
	if err != nil {
		fmt.Println("prodKeyErr")
		fmt.Println(err)
	}
	respByte, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("prodKeyErr")
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(respByte, &r)
	if err != nil {
		fmt.Println("prodKeyErr")
		fmt.Println(err)
		return
	}
	return r
}
func getCaptcha(client *req.Client, prodId string, pocCode string, scheduleNo string, chk string, key string, tYn string, sellTypeCode string) string {
	logging.Info("获取验证码图片")
	var c model.CheckCaptcha
	u := "https://tkglobal.melon.com/reservation/popup/onestop.htm"

	resp, err := client.ImpersonateChrome().R().
		// 不用手动覆盖 UA / sec-ch-ua，用 ImpersonateChrome 默认值即可
		SetHeaders(map[string]string{
			"accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
			"accept-language": "zh-CN,zh;q=0.9",
			"cache-control":   "max-age=0",
			"content-type":    "application/x-www-form-urlencoded",

			"origin":   "https://m.globalticket.melon.com",
			"priority": "u=0, i",
			"referer":  "https://m.globalticket.melon.com/",
		}).SetFormData(map[string]string{
		"chk":           url.QueryEscape(chk),
		"prodId":        prodId,
		"scheduleNo":    scheduleNo,
		"pocCode":       pocCode,
		"sellTypeCode":  sellTypeCode,
		"sellCondNo":    "1",
		"netfunnel_key": ":key=" + key + "&",
		"tYn":           tYn,
	}).
		Post(u)

	if err != nil {
		fmt.Println(err)
	}
	doc := soup.HTMLParse(resp.String())
	defer func() {
		if rs := recover(); rs != nil {
			fmt.Println(rs)
		}
	}()

	if !strings.Contains(resp.String(), "img") && strings.Contains(resp.String(), "/tktapi/glb/product/inform.json") {
		fmt.Println(resp)
		return "noCheck"
	}
	if strings.Contains(resp.String(), "403") {
		fmt.Println("403")
	}
	img := doc.Find("img")
	//fmt.Println(img.Attrs())
	input := doc.FindAll("input")

	for _, v := range input {
		if v.Attrs()["id"] == "captchaEncStr" {
			c.Chkcapt = v.Attrs()["value"]
		}

	}

	if img.Error == nil && len(img.Attrs()) > 0 {
		if len(strings.Split(img.Attrs()["src"], ",")) > 0 {
			datas := strings.Split(img.Attrs()["src"], ",")[1]
			d, _ := base64.StdEncoding.DecodeString(datas)
			fmt.Println("保存验证码")
			openFile, e := os.OpenFile("captcha.png", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
			if e != nil {
				fmt.Println(e)
			}
			defer openFile.Close()
			openFile.Write(d)
			dir, _ := os.Getwd()
			fmt.Println(dir)
			fmt.Println(url.QueryEscape(`captcha.png`))
			c.UserCaptStr = getCode(url.QueryEscape(dir + `\captcha.png`))
			//fmt.Println(url.QueryEscape(dir + `\captcha.png`))
			//c.UserCaptStr = getCode(url.QueryEscape(dir + `\captcha.png`))
			return checkCaptcha(client, prodId, pocCode, scheduleNo, c, sellTypeCode)
		}
	}

	return ""

}
func checkCaptcha(client *req.Client, prodId string, pocCode string, scheduleNo string, c model.CheckCaptcha, sellTypeCode string) string {

	var d model.CaptchaData
	u := "https://tkglobal.melon.com/reservation/ajax/checkCaptcha.json"
	params := url.Values{}
	fmt.Println("--1----")
	params.Add("userCaptStr", c.UserCaptStr)
	params.Add("chkcapt", c.Chkcapt)
	params.Add("prodId", prodId)
	params.Add("scheduleNo", scheduleNo)
	params.Add("pocCode", pocCode)
	params.Add("sellTypeCode", sellTypeCode)
	resp, err := client.ImpersonateChrome().R().
		// 不用手动覆盖 UA / sec-ch-ua，用 ImpersonateChrome 默认值即可
		SetHeaders(map[string]string{
			"accept":          "text/javascript, application/javascript, application/ecmascript, application/x-ecmascript, */*; q=0.01",
			"accept-language": "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6,ko;q=0.5",
			"cache-control":   "no-cache",
			"content-type":    "application/x-www-form-urlencoded; charset=UTF-8",

			"origin":   "https://m.globalticket.melon.com",
			"priority": "u=0, i",
			"referer":  "https://tkglobal.melon.com/reservation/popup/onestop.htm",
		}).SetFormData(map[string]string{
		"userCaptStr":  c.UserCaptStr,
		"chkcapt":      c.Chkcapt,
		"prodId":       prodId,
		"pocCode":      pocCode,
		"scheduleNo":   scheduleNo,
		"sellTypeCode": sellTypeCode,
	}).
		Post(u)
	if err != nil {
		fmt.Println(err)
	}
	//respByte, _ := http2MelonQuery("https://tkglobal.melon.com/reservation/ajax/checkCaptcha.json", params)
	respByte, err := resp.ToBytes()
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(respByte, &d)

	if d.Code == "0000" {
		logging.Info("获取验证码成功")
		fmt.Println("获取验证码成功")
		fmt.Println(d.Data)
		return d.Data
	}

	return ""
}
func getCode(fileName string) string {
	fmt.Println("自动识别验证码")
	client := &http.Client{}

	fmt.Println("http://127.0.0.1:8000/getCode/" + fileName)
	req, err := http.NewRequest("GET", "http://127.0.0.1:8000/getCode/"+fileName, nil)
	if err != nil {
		fmt.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	var c model.Code

	json.Unmarshal(bodyText, &c)

	fmt.Println(c.Code)

	return strings.ToLower(c.Code)
}
