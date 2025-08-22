package summary

import (
	"encoding/json"
	"fmt"
	"github.com/anaskhan96/soup"
	"github.com/imroc/req/v3"
	"github.com/jordan-wright/email"
	"io/ioutil"
	"log"
	"melonNew/global"
	"melonNew/httpGet"
	"melonNew/logging"
	"melonNew/model"
	"net/http"
	"net/smtp"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

func Save(c *req.Client, saveReqData model.SaveReqData) (string, error) {

	logging.Info("执行保存操作save")
	params := url.Values{}

	params.Add("jType", "I")
	params.Add("delvyTypeCode", "DV0002")
	fmt.Println(global.GlobalConfig.GetString("userName"))
	fmt.Println(global.GlobalConfig.GetString("email"))
	params.Add("userName", global.GlobalConfig.GetString("userName"))
	params.Add("tel", global.GlobalConfig.GetString("tel"))
	params.Add("email", global.GlobalConfig.GetString("email"))
	params.Add("rd2", "2")
	params.Add("delvyCategory", "0")
	params.Add("destntSeq", "")
	params.Add("destntName", "")
	params.Add("bassDestntYn", "N")
	params.Add("delvyName", "")
	params.Add("delvyTel", "")
	params.Add("delvyZipno", "")
	params.Add("delvyAddrDtl", "")
	params.Add("delvyName", "")
	params.Add("recv_country", "")
	params.Add("recv_name", "")
	params.Add("recv_address", "")

	params.Add("recv_city", "")
	params.Add("recv_state", "")
	params.Add("recv_zipno", "")
	params.Add("recv_tel1", "")
	params.Add("recv_tel2", "")
	params.Add("addAddress", "")
	params.Add("recv_country_code", "")
	params.Add("recv_delvy_price", "0")
	params.Add("payMethodCode", "AP0012") //AP0010 支付宝 AP0012 信用卡
	params.Add("cardCode", "FOREIGN_CHINABANK")
	params.Add("cardCodeName", "UnionPay")
	params.Add("autheTypeCode", "AT0005")
	params.Add("cardQuota", "12")
	params.Add("quota", "00")
	params.Add("cashReceiptIssueCode", "0")
	params.Add("cashReceiptRegType", "0")
	params.Add("cashReceiptRegType2", "10")
	params.Add("cashReceiptRegTelNo1", "010")
	params.Add("chkAgreeAll", "on")
	params.Add("chkAgree", "on")
	params.Add("prodId", saveReqData.ProdId)
	params.Add("pocCode", saveReqData.PocCode)
	params.Add("scheduleNo", saveReqData.ScheduleNo)
	//params.Add("netfunnel_key", "")
	params.Add("rsrvVolume", "1")
	params.Add("payAmt", strconv.Itoa(saveReqData.PayAmt))
	params.Add("priceNo", saveReqData.PriceNo)
	params.Add("seatId", saveReqData.SeatId)

	//seatNm 中L与R从 座位表中取
	//params.Add("seatInfoListWithPriceType", `[{"priceNo":10067,"seatId":`+saveReqData.SeatId+`,"clipSeatId":null,"gradeNm":"지정석","seatNm":"  ［A．B GATE］402 구역 `+saveReqData.R+` 열 `+saveReqData.S+` 번 ","basePrice":110000,"priceName":"기본가","sejongPriceCode":null}]`)
	seatInfo := saveReqData.FloorNo + "_" + saveReqData.FloorName + "_" + saveReqData.AreaNo + "_" + saveReqData.AreaName + "_" + saveReqData.R + "_" + saveReqData.S + "_"

	params.Add("seatInfoListWithPriceType", `[{"priceNo":`+saveReqData.PriceNo+`,"seatId":"`+saveReqData.SeatId+`","clipSeatId":null,"gradeNm":"`+saveReqData.SeatGradeName+`","seatNm":"`+saveReqData.FloorNo+` `+saveReqData.FloorName+` `+saveReqData.AreaNo+` `+saveReqData.AreaName+` `+saveReqData.R+` 열 `+saveReqData.S+` 번","basePrice":`+saveReqData.BasePrice+`,"priceName":"`+saveReqData.PriceName+`","sejongPriceCode":null}]`)

	params.Add("firstSeatId", saveReqData.SeatId)
	params.Add("sellTypeCode", saveReqData.SellTypeCode)
	params.Add("chkcapt", saveReqData.Chkcapt)
	//body := strings.NewReader(params.Encode())
	//st := `addAddress=&autheTypeCode=AT0005&bassDestntYn=N&cardCode=FOREIGN_CHINABANK&cardCodeName=UnionPay&cardQuota=12&cashReceiptIssueCode=0&cashReceiptRegTelNo1=010&cashReceiptRegType=0&cashReceiptRegType2=10&chkAgree=on&chkAgreeAll=on&chkcapt=PuCT8XWLErUxtLhH9v%2FypS1McGNmzrwueYyrIj6Nmws%3D&delvyAddrDtl=&delvyCategory=0&delvyName=&delvyName=&delvyTel=&delvyTypeCode=DV0002&delvyZipno=&destntName=&destntSeq=&email=2055663800%40qq.com&firstSeatId=598_1580&jType=I&payAmt=146000&payMethodCode=AP0012&pocCode=SC0002&priceNo=10067&prodId=210230&quota=00&rd2=2&recv_address=&recv_city=&recv_country=&recv_delvy_price=0&recv_name=&recv_state%3A+=&recv_tel1=&recv_tel2=&recv_zipno=&rsrvVolume=1&scheduleNo=100002&seatId=598_1580&seatInfoListWithPriceType=%5B%7B%22priceNo%22%3A10067%2C%22seatId%22%3A598_1580%2C%22clipSeatId%22%3Anull%2C%22gradeNm%22%3A%22%ED%98%84%EC%9E%A5%EC%88%98%EB%A0%B9S%EC%84%9D%282%ED%9A%8C%EC%B0%A8%29%22%2C%22seatNm%22%3A%22%EC%84%9C9%EB%AC%B8+1+%EC%B8%B5+%EC%84%9C%EC%B8%A1+G+%EA%B5%AC%EC%97%AD+33+%EC%97%B4+5+%EB%B2%88%22%2C%22basePrice%22%3A143000%2C%22priceName%22%3A%22%EA%B8%B0%EB%B3%B8%EA%B0%80%22%2C%22sejongPriceCode%22%3Anull%7D%5D&sellTypeCode=ST0001&tel=13024516491&userName=XUMUXI`
	//body = strings.NewReader(st)
	fmt.Println("请求save")
	strMap := make(map[string]string, len(params))
	for key, values := range params {
		if len(values) > 0 {
			strMap[key] = values[0]
		}
	}
	resp, err := httpGet.HttpMelonQueryPost(c, "https://tkglobal.melon.com/tktapi/glb/reservation/save.json?v=1", strMap)
	if err != nil {
		fmt.Println(err)
	}
	var data model.SaveHandler
	json.Unmarshal(resp, &data)
	//str := `{"flplanTypeCode":"DR0002","code":"0000","seatInfoListWithPriceType":"[{\"priceNo\":10067,\"seatId\":\"602_224\",\"clipSeatId\":null,\"gradeNm\":\"일반석\",\"seatNm\":\"2 층 39 구역 11 열 5 번 \",\"basePrice\":121000,\"priceName\":\"기본가\",\"sejongPriceCode\":null}]","cardCode":"FOREIGN_CHINABANK","jtype":"I","eType":"","cust_ip":"34.96.5.73","prodId":"210031","kakaoPayType":"","userName":"LIU JIAHUI","payAmt":"126000","perfMainName":"2024 RIIZE FAN－CON ‘RIIZING DAY’ FINALE in SEOUL","payNo":"2024071910731969","midOptionKey":"온라인_해외_인증_일반","staticDomain":null,"httpsDomain":null,"quota":"00","tel":"17705051327","rsrvSeq":"2024071907703891","payMethodCode":"AP0012","httpDomain":null}`
	//saveReqData.SeatId = "602_224"

	//json.Unmarshal([]byte(str), &data)
	if len(data.RsrvSeq) > 0 {
		fmt.Println("success")
		//data.RsrvSeq = "2024031105158306"
		//data.PayNo = "2024031107587553"
		logging.Info(saveReqData.ScheduleNo)
		logging.Info(params.Encode())
		logging.Info("保存时间" + time.Now().String())
		fmt.Println("保存时间" + time.Now().String())
		s := payInitForm(c, data, seatInfo)

		return s, nil
	} else {
		fmt.Println("-11---1-1ssss-1-1--")
		fmt.Println(string(resp))
		if strings.Contains(string(resp), "결제 금액 오류") {
			logging.Error(resp)
		}
		logging.Error(string(resp))
		logging.Info(params.Encode())

		return "error", fmt.Errorf(string(resp))
	}

}

func payInitForm(c *req.Client, d model.SaveHandler, seatId string) string {
	logging.Info("提交表单....")
	params := url.Values{}

	params.Add("flplanTypeCode", d.FlplanTypeCode)
	params.Add("code", "0000")
	params.Add("seatInfoListWithPriceType", d.SeatInfoListWithPriceType)
	params.Add("cardCode", d.CardCode)
	params.Add("jtype", "I")
	params.Add("eType", "")

	params.Add("cust_ip", d.CustIP)
	params.Add("prodId", d.ProdID)
	params.Add("kakaoPayType", "")
	params.Add("userName", d.UserName)
	params.Add("payAmt", d.PayAmt)
	params.Add("perfMainName", d.PerfMainName)

	params.Add("payNo", d.PayNo)
	params.Add("midOptionKey", d.MidOptionKey)
	params.Add("staticDomain", "")
	params.Add("httpsDomain", "")
	params.Add("quota", d.Quota)
	params.Add("tel", d.Tel)
	params.Add("rsrvSeq", d.RsrvSeq)

	params.Add("payMethodCode", d.PayMethodCode)
	params.Add("httpDomain", "")
	params.Add("card_pay_method", "GLB")
	strMap := make(map[string]string, len(params))
	for key, values := range params {
		if len(values) > 0 {
			strMap[key] = values[0]
		}
	}
	//body := strings.NewReader(params.Encode())
	//var body = strings.NewReader(`flplanTypeCode=DR0002&code=0000&seatInfoListWithPriceType=%5B%7B%22priceNo%22%3A10067%2C%22seatId%22%3A%22188_393%22%2C%22clipSeatId%22%3Anull%2C%22gradeNm%22%3A%22%EC%A7%80%EC%A0%95%EC%84%9D%22%2C%22seatNm%22%3A%22%EF%BC%BBA%EF%BC%8EB+GATE%EF%BC%BD402+%EA%B5%AC%EC%97%AD+L+%EC%97%B4+37+%EB%B2%88+%22%2C%22basePrice%22%3A110000%2C%22priceName%22%3A%22%EA%B8%B0%EB%B3%B8%EA%B0%80%22%2C%22sejongPriceCode%22%3Anull%7D%5D&cardCode=FOREIGN_CHINABANK&jtype=I&eType=&cust_ip=218.250.158.239&prodId=209540&kakaoPayType=&userName=jili+zheng&payAmt=112000&perfMainName=2024+EXO+FAN+MEETING%EF%BC%9AONE&payNo=2024031207607649&midOptionKey=%EC%98%A8%EB%9D%BC%EC%9D%B8_%ED%95%B4%EC%99%B8_%EC%9D%B8%EC%A6%9D_%EC%9D%BC%EB%B0%98&staticDomain=&httpsDomain=&quota=00&tel=13055253090&rsrvSeq=2024031205174677&payMethodCode=AP0001&httpDomain=&card_pay_method=GLB`)
	resp, err := httpGet.HttpMelonQueryPost(c, "https://tkglobal.melon.com/reservation/ajax/payInitForm.htm?procMode=R", strMap)
	if err != nil {
		logging.Info(err)
	}

	doc := soup.HTMLParse(string(resp))

	src := doc.Find("iframe").Attrs()["src"]
	u, err := url.Parse(src)
	if err != nil {
		log.Fatal(err)
	}

	queryParams := u.Query()

	if len(queryParams["session_key"]) > 0 {
		sessionKey := queryParams["session_key"][0]
		logging.Info("sessionKey为" + sessionKey)
		fmt.Println(queryParams)

		location1 := payFormFirst(sessionKey, queryParams["dkpg_payment_id"][0])
		fmt.Println(location1)
		location2, session := userAuthForm(location1)
		fmt.Println(location2)
		location3 := userAuthFormR(location2, session)
		fmt.Println(location3)
		sessionKey1, path := payForm(location3)
		if sessionKey1 != "" {
			emails(seatId, path)
			return "success"
		} else {
			return "fail"
		}

	}
	return ""
}
func emails(u1, u2 string) {
	e := email.NewEmail()
	dir, _ := os.Getwd()
	//设置发送方的邮箱
	e.From = "dj <736411153@qq.com>"
	// 设置接收方的邮箱
	e.To = []string{"283384082@qq.com"}
	//设置主题
	e.Subject = "Awesome web"
	e.Text = []byte("Text Body is, of course, supported!" + dir + "\n" + u1 + "\n" + "------------------------------" + "\n" + u2)
	//这块是设置附件
	//e.AttachFile("./test.txt")
	//设置服务器相关的配置
	err := e.Send("smtp.qq.com:25", smtp.PlainAuth("", "736411153@qq.com", "xndtagtgdxvdbffc", "smtp.qq.com"))
	if err != nil {
		fmt.Println(err)
	}
}
func payFormFirst(sessionKey string, dkpgPaymentId string) string {
	logging.Info("payFormFirst")
	var client *http.Client
	if global.GlobalConfig.GetString("openProxy") == "1" {
		proxy := global.GlobalConfig.GetString("proxy")
		proxyAddr, _ := url.Parse(proxy)
		client = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyAddr),
			},
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse /* 不进入重定向 */
			},
		}
	} else {

		client = &http.Client{
			Transport: &http.Transport{},
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse /* 不进入重定向 */
			},
		}
	}

	//u:=`https://dkpg-web.payments.kakao.com/dkpg/v1/pay_form?&dkpg_payment_id=8030000354505630&session_key=1724284416287-9f126111-1d6f-46a5-a8ae-1d1bf0ce6066`
	req, err := http.NewRequest("GET", "https://dkpg-web.payments.kakao.com/dkpg/v1/pay_form?&dkpg_payment_id="+dkpgPaymentId+"&session_key="+sessionKey, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Referer", "https://tkglobal.melon.com/")
	req.Header.Set("Sec-Fetch-Dest", "iframe")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36 Edg/122.0.0.0")
	req.Header.Set("sec-ch-ua", `"Chromium";v="122", "Not(A:Brand";v="24", "Microsoft Edge";v="122"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	resp, err := client.Do(req)

	if err != nil {

		logging.Info(err)
	}

	return resp.Header.Get("location")

}
func userAuthForm(path string) (string, string) {
	logging.Info("userAuthForm")

	var client *http.Client
	if global.GlobalConfig.GetString("openProxy") == "1" {
		proxy := global.GlobalConfig.GetString("proxy")
		proxyAddr, _ := url.Parse(proxy)
		client = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyAddr),
			},
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse /* 不进入重定向 */
			},
		}
	} else {

		client = &http.Client{
			Transport: &http.Transport{},
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse /* 不进入重定向 */
			},
		}
	}
	session := strings.Split(path, "session_key=")[1]
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", "_sweb_session="+session)
	req.Header.Set("Host", "dkpg-web.payments.kakao.com")
	req.Header.Set("Referer", "https://tkglobal.melon.com/")
	req.Header.Set("Sec-Fetch-Dest", "iframe")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36 Edg/122.0.0.0")
	req.Header.Set("sec-ch-ua", `"Chromium";v="122", "Not(A:Brand";v="24", "Microsoft Edge";v="122"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	resp, err := client.Do(req)
	if err != nil {
		logging.Info(err)
	}
	defer resp.Body.Close()
	return resp.Header.Get("location"), session

}
func userAuthFormR(path string, session string) string {
	logging.Info("userAuthFormR")
	var client *http.Client
	if global.GlobalConfig.GetString("openProxy") == "1" {
		proxy := global.GlobalConfig.GetString("proxy")
		proxyAddr, _ := url.Parse(proxy)
		client = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyAddr),
			},
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse /* 不进入重定向 */
			},
		}
	} else {

		client = &http.Client{
			Transport: &http.Transport{},
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse /* 不进入重定向 */
			},
		}
	}
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Connection", "keep-alive")

	req.Header.Set("Cookie", "_sweb_session="+session)
	req.Header.Set("Host", "dkpg-web.payments.kakao.com")
	req.Header.Set("Referer", "https://tkglobal.melon.com/")
	req.Header.Set("Sec-Fetch-Dest", "iframe")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36 Edg/122.0.0.0")
	req.Header.Set("sec-ch-ua", `"Chromium";v="122", "Not(A:Brand";v="24", "Microsoft Edge";v="122"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	return resp.Header.Get("location")
}
func payForm(path string) (string, string) {
	logging.Info("payForm")
	u, _ := url.Parse(path)
	sessionKey := u.Query()["session_key"][0]
	var client *http.Client
	if global.GlobalConfig.GetString("openProxy") == "1" {
		proxy := global.GlobalConfig.GetString("proxy")
		proxyAddr, _ := url.Parse(proxy)
		client = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyAddr),
			},
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse /* 不进入重定向 */
			},
		}
	} else {

		client = &http.Client{
			Transport: &http.Transport{},
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse /* 不进入重定向 */
			},
		}
	}
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", "_sweb_session="+sessionKey)
	req.Header.Set("Referer", "https://tkglobal.melon.com/")
	req.Header.Set("Sec-Fetch-Dest", "iframe")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36 Edg/122.0.0.0")
	req.Header.Set("sec-ch-ua", `"Chromium";v="122", "Not(A:Brand";v="24", "Microsoft Edge";v="122"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var seesionKey string
	doc := soup.HTMLParse(string(bodyText))
	src := doc.FindAll("input")

	for k, v := range src {
		if v.Attrs()["name"] == "session_key" {

			seesionKey = src[k].Attrs()["value"]
		}
	}
	return seesionKey, path

}
