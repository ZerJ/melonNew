package ticket

import (
	"fmt"
	"github.com/imroc/req/v3"
	"log"
	"melonNew/global"
	"melonNew/logging"
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
	var keyData model.ProdKey
	var wkey, chkcapt string
	// 定时开始
	ts := global.GlobalConfig.GetString("ts")
	if ts == "" {
		ts = "2025-08-22 18:59:59.7312"
	}
	startTime, _ := time.ParseInLocation("2006-01-02 15:04:05", ts, time.Local)
	for time.Now().Before(startTime) {
		time.Sleep(2 * time.Millisecond)
	}
	for {

		if time.Now().Local().After(startTime) {
			logging.Info("start:", time.Now())
			keyData = prodKey(c, prodId, scheduleNos[0])
			if len(keyData.Key) > 0 {
				break
			}

		}

		//chk, trafficCtrlYn, nflActId = prodKey(prodId, scheduleNo[0])
		//if len(chk) > 0 {
		//	break
		//}

	}
	// 1. 获取产品 Key

	if keyData.Key == "" {
		log.Fatal("prodKey 获取失败")
	}
	// 2. 获取 wseq key
	if len(keyData.NflActID) > 0 {
		WSeqkey, r, err := WSeqAuto(c, keyData.NflActID, 0)

		if err != nil {
			fmt.Println(err)
		}
		c.SetCommonCookies(&http.Cookie{Name: "NetFunnel_ID", Value: r, Domain: ".melon.com"})
		global.GlobalConfig.Set("key", WSeqkey)
		wkey = WSeqkey
	}
	// 2. 获取 wseq key

	//
	//// 3. 获取验证码

	for {
		chkcapt = getCaptcha(c, prodId, pocCode, scheduleNos[0], keyData.Key, wkey, keyData.TrafficCtrlYn, sellTypeCode)
		fmt.Println("--1-chat")
		if chkcapt != "" {
			break
		}
		time.Sleep(5 * time.Second)
	}
	if len(keyData.NflActID) > 0 {
		go ws(c, keyData.NflActID)
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
func ws(c *req.Client, nflActId string) {
	for {
		time.Sleep(500 * time.Second)
		if len(nflActId) > 0 {
			WSeqkey, r, err := WSeqAuto(c, nflActId, 0)
			if err != nil {
				fmt.Println(err)
			}
			//key = ":key=380480F642DE9C1DF3448378031F377CDE96E4982C8F183D67451A3ACEDF8B8A534C84197F09953206DD23138F2C4EDE14EE497C41731C2C8C2376F89B7B2BC6591AC788FB745524F430A8F20D6BCB249E00242B2D654AA6069D45CB624390822E18C91D0C404445A11B9A85FDC944B32C382C312C34323732312C30"
			setCookie(r)
			global.GlobalConfig.Set("key", WSeqkey)
			//fmt.Println(key)
		}
	}

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
		time.Sleep(1 * time.Second)
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
