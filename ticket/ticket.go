package ticket

import (
	"fmt"
	"github.com/imroc/req/v3"
	"log"
	"melonNew/global"
	"melonNew/model"
	"melonNew/summary"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

func GetTicket(c *req.Client) {
	prodId := global.GlobalConfig.GetString("prodId")
	sellTypeCode := global.GlobalConfig.GetString("sellTypeCode")
	scheduleNos := strings.Split(global.GlobalConfig.GetString("scheduleNo"), ",")
	pocCode := "SC0002"
	var WSeqkey, chkcapt string
	// 定时开始
	ts := global.GlobalConfig.GetString("ts")
	if ts == "" {
		ts = "2025-08-20 18:59:59"
	}
	startTime, _ := time.ParseInLocation("2006-01-02 15:04:05", ts, time.Local)
	for time.Now().Before(startTime) {
		time.Sleep(200 * time.Millisecond)
	}

	// 1. 获取产品 Key
	keyData := prodKey(c, prodId, scheduleNos[0])

	if keyData.Key == "" {
		log.Fatal("prodKey 获取失败")
	}
	// 2. 获取 wseq key
	if len(keyData.NflActID) > 0 {
		WSeqkey, r, err := WSeqAuto(c, keyData.NflActID, 5)
		if err != nil {
			fmt.Println(err)
		}
		c.SetCommonCookies(&http.Cookie{Name: "NetFunnel_ID", Value: r, Domain: ".melon.com"})
		global.GlobalConfig.Set("key", WSeqkey)
	}
	// 2. 获取 wseq key

	//
	//// 3. 获取验证码

	for {
		chkcapt = getCaptcha(c, prodId, pocCode, scheduleNos[0], keyData.Key, WSeqkey, keyData.TrafficCtrlYn, sellTypeCode)
		if chkcapt == "noCheck" || chkcapt != "" {
			break
		}
		time.Sleep(5 * time.Second)
	}
	global.GlobalConfig.Set("chkcapt", chkcapt)
	//
	// 4. 开 goroutine 抢票
	var wg sync.WaitGroup
	go getChkcapt(c, prodId, pocCode, scheduleNos[0], keyData.Key, sellTypeCode)
	for _, v := range scheduleNos {
		i := summary.InformProdSch(c, prodId, pocCode, v, sellTypeCode)
		d, amap := summary.S1(c, prodId, pocCode, v, i.ProdInform.PerfStartDay)
		block1, _ := summary.Summary(c, prodId, pocCode, v, i.ProdInform.PerfStartDay, d, amap)

		wg.Add(1)
		go work(c, v, i, sellTypeCode, block1)
	}
	wg.Wait()
}
func work(c *req.Client, scheduleNo string, i model.PromReqData, sellTypeCode string, block1 []model.SeatReq) {
	fmt.Println("开始处理场次:", scheduleNo)

	for {
		chkcapt := global.GlobalConfig.GetString("chkcapt")
		key := global.GlobalConfig.GetString("key")

		if len(block1) == 0 || chkcapt == "" {
			time.Sleep(time.Second)
			continue
		}

		if chkcapt == "noCheck" {
			chkcapt = ""
		}

		success := false

		for _, v := range block1 {
			if summary.SeatWork(c, v, i, chkcapt, key, sellTypeCode, block1) {
				success = true
			}
			time.Sleep(2 * time.Second) // 防止过快请求
		}

		if success {
			fmt.Println("----")
		}

		fmt.Println("本轮未成功，继续重试...")
		time.Sleep(time.Second)
	}

	fmt.Println("完成场次:", scheduleNo)
}

func setCookie(str string) {
	var cookie string
	var j string
	//if strings.Contains(str, "5002:201") {
	//	str = strings.Replace(str, "5002:201", "5002:200", 1)
	//}
	for _, b := range strings.Split(global.GlobalConfig.GetString("melonCookie"), ";") {

		if strings.Split(b, "=")[0] == " NetFunnel_ID" {

			j = strings.Split(b, "=")[0] + "=" + url.QueryEscape(str)
		} else {
			j = b
		}
		cookie = cookie + "" + j + ";"

	}
	if !strings.Contains(cookie, "NetFunnel_ID") {
		cookie = cookie + "; NetFunnel_ID=" + url.QueryEscape(str)
	}

	global.GlobalConfig.Set("melonCookie", strings.TrimRight(cookie, ";"))
	fmt.Println(global.GlobalConfig.GetString("melonCookie"))

	//return strings.Split(strings.Split(str, ":200")[1], "&nwait=")[0]
}
func getChkcapt(c *req.Client, prodId string, pocCode string, scheduleNo string, chk string, sellTypeCode string) {
	for {

		time.Sleep(400 * time.Second)
		for {
			chkcapt := getCaptcha(c, prodId, pocCode, scheduleNo, chk, "", "", sellTypeCode)

			if len(chkcapt) > 0 && chkcapt != "noCheck" {
				global.GlobalConfig.Set("chkcapt", chkcapt)
				break
			} else if chkcapt == "noCheck" {
				chkcapt = "noCheck"
				global.GlobalConfig.Set("chkcapt", chkcapt)
				break
			}
			time.Sleep(20 * time.Second)
		}

	}

}
