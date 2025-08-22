package summary

import (
	"encoding/json"
	"fmt"
	"github.com/imroc/req/v3"
	"melonNew/httpGet"
	"melonNew/logging"
	"melonNew/model"
	"net/url"
	"strconv"
	"strings"
)

func SeatWork(
	c *req.Client,
	block model.SeatReq,
	i model.PromReqData,
	chkcapt string,
	key string,
	sellTypeCode string,
	blocks []model.SeatReq,
) bool {
	r, err := prodLimit(c, i, block, chkcapt, key, sellTypeCode, blocks)
	if err != nil {
		fmt.Println("SeatWork 错误:", err)
		return false
	}

	if r == "success" {
		fmt.Println("SeatWork 成功:", block.BlockID)
		return true
	}

	return false
}

func prodLimit(c *req.Client, promReq model.PromReqData, block model.SeatReq, chk string, key string, sellTypeCode string, blocks []model.SeatReq) (r string, err error) {

	//fmt.Println(block)
	var saveData model.SaveReqData
	var data model.ProdLimitData
	seat, err := GetSeatList(c, block)
	//fmt.Println(time.Now())
	for _, v := range seat.SeatData.St {

		for _, b := range v.Ss {

			if len(b.Sid) > 0 {
				fmt.Println(b.Sid)
				fmt.Println("-------")
				logging.Info("开始执行买票操作1：prodlimit"+block.BlockID, b.Sid)
				//
				//d := strings.Split(b.Sid, "_")[1]
				//idInt, _ := strconv.Atoi(d)
				//if idInt > 140 {
				//	continue
				//}
				data.SeatIds = b.Sid
				for k, ss := range blocks {
					if ss.SeatGradeName == b.Gn && ss.BlockID == block.BlockID {
						block = blocks[k]
					}
				}
				saveData.R = b.Rn
				saveData.S = b.Sn

				saveData.Chkcapt = chk

				//saveData.PriceNo = "10067"
				fmt.Println(b.Gn)

				saveData.SeatGradeName = b.Gn

				saveData.PriceName = "기본가"
				payloadMap := map[string]string{
					"langCd":           "EN",
					"prodId":           block.ProdID,
					"pocCode":          "SC0002",
					"perfTypeCode":     "GN0001",
					"perfDate":         block.PerfStartDay,
					"scheduleNo":       block.ScheduleNo,
					"sellTypeCode":     sellTypeCode,
					"sellCondNo":       "",
					"perfMainName":     promReq.ProdInform.PerfMainName,
					"seatGradeNo":      block.SeatGradeNo,
					"seatGradeName":    block.SeatGradeName,
					"blockId":          block.BlockID,
					"sntv":             url.QueryEscape(block.Sntv),
					"blockTypeCode":    "",
					"floorNo":          block.FloorNo,
					"floorName":        block.FloorName,
					"areaNo":           block.AreaNo,
					"areaName":         block.AreaName,
					"prodTypeCode":     promReq.ProdInform.ProdTypeCode,
					"flplanTypeCode":   "DR0002",
					"scheduleTypeCode": "SG0001",
					"seatTypeCode":     "SE0001",
					"jType":            "I",
					"cardGroupId":      "",
					"cardBpId":         "",
					"cardMid":          "",
					"rsrvStep":         "SAT",
					"zamEnabled":       "0",
					"zamKey":           "",
					"trafficCtrlYn":    "N",
					"netfunnel_key":    key,
					"stvn_view_list":   url.QueryEscape(block.StvnViewList),
					"mapClickYn":       "Y",
					"seatId":           data.SeatIds,
					"clipSeatId":       "",
					"chkcapt":          url.QueryEscape(chk),
				}
				logging.Info(`langCd=EN&prodId=` + block.ProdID + `&pocCode=SC0002&perfTypeCode=GN0001&perfDate=` + block.PerfStartDay + `&scheduleNo=` + block.ScheduleNo + `&sellTypeCode=` + sellTypeCode + `&sellCondNo=&perfMainName=` + url.QueryEscape(promReq.ProdInform.PerfMainName) + `&seatGradeNo=` + block.SeatGradeNo + `&seatGradeName=` + block.SeatGradeName + `&blockId=` + block.BlockID + `&sntv=` + url.QueryEscape(block.Sntv) + `&blockTypeCode=&floorNo=` + block.FloorNo + `&floorName=` + url.QueryEscape(block.FloorName) + `&areaNo=` + block.AreaNo + `&areaName=` + url.QueryEscape(block.AreaName) + `&prodTypeCode=` + promReq.ProdInform.ProdTypeCode + `&flplanTypeCode=DR0002&scheduleTypeCode=SG0001&seatTypeCode=SE0001&jType=I&cardGroupId=&cardBpId=&cardMid=&rsrvStep=SAT&zamEnabled=0&zamKey=&trafficCtrlYn=N&netfunnel_key=` + url.QueryEscape(key) + `&stvn_view_list=` + url.QueryEscape(block.StvnViewList) + `&mapClickYn=Y&seatId=` + data.SeatIds + `&clipSeatId=&chkcapt=` + url.QueryEscape(chk))
				//payload := strings.NewReader(`langCd=EN&prodId=` + block.ProdID + `&pocCode=SC0002&perfTypeCode=GN0001&perfDate=` + block.PerfStartDay + `&scheduleNo=` + block.ScheduleNo + `&sellTypeCode=` + sellTypeCode + `&sellCondNo=&perfMainName=` + url.QueryEscape(promReq.ProdInform.PerfMainName) + `&seatGradeNo=` + block.SeatGradeNo + `&seatGradeName=` + block.SeatGradeName + `&blockId=` + block.BlockID + `&sntv=` + url.QueryEscape(block.Sntv) + `&blockTypeCode=&floorNo=` + block.FloorNo + `&floorName=` + url.QueryEscape(block.FloorName) + `&areaNo=` + block.AreaNo + `&areaName=` + url.QueryEscape(block.AreaName) + `&prodTypeCode=` + promReq.ProdInform.ProdTypeCode + `&flplanTypeCode=DR0002&scheduleTypeCode=SG0001&seatTypeCode=SE0001&jType=I&cardGroupId=&cardBpId=&cardMid=&rsrvStep=SAT&zamEnabled=0&zamKey=&trafficCtrlYn=N&netfunnel_key=` + url.QueryEscape(key) + `&stvn_view_list=` + url.QueryEscape(block.StvnViewList) + `&mapClickYn=Y&seatId=` + data.SeatIds + `&clipSeatId=&chkcapt=2IxvFGPe9BHo6rNbVBSQlnc5Ib%2F4uOWjC9T7qxFbQvQ%3D`)
				//var data1 = strings.NewReader(`langCd=EN&prodId=209657&pocCode=SC0002&perfTypeCode=GN0001&perfDate=` + block.PerfStartDay + `&scheduleNo=100001&sellTypeCode=ST0002&sellCondNo=1&perfMainName=2024+NCT+DREAM+WORLD+TOUR+%E3%80%88THE+DREAM+SHOW+3%EF%BC%9ADREAM%EF%BC%88%EF%BC%89SCAPE%E3%80%89&seatGradeNo=&seatGradeName=&blockId=` + block.BlockID + `&sntv=` + block.Sntv + `&blockTypeCode=&floorNo=&floorName=&areaNo=` + block.AreaNo + `&areaName=` + url.QueryEscape(block.AreaName) + `&prodTypeCode=PT0001&flplanTypeCode=DR0002&scheduleTypeCode=SG0001&seatTypeCode=SE0001&jType=I&cardGroupId=&cardBpId=&cardMid=&rsrvStep=SAT&zamEnabled=0&zamKey=&trafficCtrlYn=Y&netfunnel_key=%3Akey%3D14D08FDF032363695953B1EE52003A88C098D9256FBB0D122FA34A419A61B8B2C51E63206D54C1511DC43E2C33D01701664F513F0B36D2197218AEA0BFA08CCEE2FF71A65C04DA1951056A0C151271F984A5D2AAE4178E8041EE17D0C18F403253FC3C73B8D45B026BF7195A271CE63D246276E72E3A6B1A5CE877B7468289082C30%26amp%3B&stvn_view_list=%2CFLOOR+F10%3B%2CFLOOR+F11%3B%2CFLOOR+F3%3B%2CFLOOR+F4%3B%2CFLOOR+F8%3B%2CFLOOR+F9%3B%2C101%3B%2C102%3B%2C103%3B%2C104%3B%2C105%3B%2C106%3B%2C107%3B%2C108%3B%2C109%3B%2C110%3B%2C111%3B%2C112%3B%2C113%3B%2C114%3B%2C201%3B%2C202%3B%2C203%3B%2C204%3B%2C205%3B%2C206%3B%2C207%3B%2C208%3B%2C209%3B%2C210%3B%2C301%3B%2C302%3B%2C303%3B%2C304%3B%2C305%3B%2C306%3B%2C307%3B%2C308%3B%2C309%3B%2C310%3B%2C311%3B%2C312%3B%2C313%3B%2C314%3B%2C315%3B%2C316%3B%2C317%3B%2C318%3B%2C319%3B%2C320%3B%2C321%3B%2C322%3B%2C401%3B%2C402%3B%2C403%3B%2C404%3B%2C405%3B%2C406%3B%2C407%3B%2C408%3B%2C409%3B%2C410%3B%2C411%3B%2C412%3B%2C413%3B%2C414%3B%2C415%3B%2C416%3B%2C417%3B%2C418%3B%2C419%3B%2C420%3B%2C421%3B%2C422%3B%2C423%3B%2C424%3B%2CD01%3B%2CD02%3B%2CD06%3B%2CD07%3B%2CFLOOR+F1%3B%2CFLOOR+F10%3B%2CFLOOR+F11%3B%2CFLOOR+F12%3B%2CFLOOR+F13%3B%2CFLOOR+F14%3B%2CFLOOR+F15%3B%2CFLOOR+F16%3B%2CFLOOR+F2%3B%2CFLOOR+F5%3B%2CFLOOR+F6%3B%2CFLOOR+F7%3B%2CFLOOR+F8%3B%2CFLOOR+F9%3B%2CT01%3B%2CT02%3B%2CT03%3B%2CT04%3B%2CT05%3B%2CT06%3B%2CT07%3B%2CT11%3B%2CT12%3B%2CT13%3B%2CT15%3B%2CT16%3B%2CT17&mapClickYn=Y&seatId=` + data.SeatIds + `&clipSeatId=&chkcapt=` + chk)
				//body := strings.NewReader(params.Encode())
				fmt.Println(data.SeatIds)
				logging.Info("准备购买block为" + block.BlockID + "的seatId 为:" + data.SeatIds)
				//fmt.Println(body)
				resp, _ := httpGet.HttpMelonQueryPost(c, "https://tkglobal.melon.com/tktapi/glb/reservation/prodlimit.json?v=1", payloadMap)
				//logging.Info(resp)
				err = json.Unmarshal(resp, &data)
				if err != nil {
					logging.Info(payloadMap)
					logging.Info(string(resp))
					fmt.Println("----")
					fmt.Println("----")
					fmt.Println(string(resp))
					return "error", err
				}

				saveData.ProdId = strconv.Itoa(promReq.ProdInform.ProdID)
				saveData.PocCode = promReq.ProdInform.PocCode
				saveData.ScheduleNo = strconv.Itoa(promReq.ScheduleNo)
				saveData.SeatId = data.SeatIds
				saveData.FloorNo = block.FloorNo

				saveData.FloorName = block.FloorName
				saveData.AreaNo = block.AreaNo
				saveData.AreaName = block.AreaName

				saveData.BasePrice = block.BasePrice
				fmt.Println(block.BasePrice)
				logging.Info(block.BasePrice)
				ps, err := strconv.Atoi(block.BasePrice)
				if err != nil {
					logging.Error(err)
				}
				logging.Info(ps)
				logging.Info(promReq.ProdInform.RsrvFee)
				//price, _ := strconv.Atoi(global.GlobalConfig.GetString("price"))
				saveData.PayAmt = ps + promReq.ProdInform.RsrvFee
				logging.Info(saveData.PayAmt)
				//saveData.PayAmt = p + price
				saveData.SellTypeCode = sellTypeCode
				//saveData.PayAmt = p + 4000

				fmt.Println("---zaz")
				if len(data.EncryptedSeatIds) > 0 {

					tickets, err := ticketType(c, promReq, data.SeatIds, block)
					logging.Info(tickets.PriceNo)
					saveData.PriceNo = tickets.PriceNo

					fmt.Println(data.SeatIds)
					fmt.Println(data.EncryptedSeatIds)
					stepTicket(c, strconv.Itoa(promReq.ProdInform.ProdID), block.ScheduleNo, data)
					if saveData.PriceNo == "0" {
						saveData.PriceNo = "14554"
					}

					r, err := Save(c, saveData)
					if err != nil {

						fmt.Println(err)
						logging.Info(err)
						continue

					}

					return r, nil
				} else {
					logging.Info(string(resp))

				}
			}
		}
	}

	return "error", nil
	//test2(data)
}
func ticketType(c *req.Client, i model.PromReqData, seatId string, block model.SeatReq) (saveDataReq model.SaveReqData, err error) {
	fmt.Println("执行ticketType")
	logging.Info("执行ticketType")
	params := url.Values{}
	params.Add("langCd", "EN")
	params.Add("prodId", strconv.Itoa(i.ProdInform.ProdID))
	params.Add("pocCode", i.ProdInform.PocCode)
	params.Add("perfTypeCode", i.ProdInform.ProdTypeCode)

	params.Add("perfDate", i.ProdInform.PerfStartDay)
	params.Add("scheduleNo", strconv.Itoa(i.ScheduleNo))
	params.Add("sellTypeCode", i.SchList[0].SellTypeCode)
	params.Add("perfMainName", i.ProdInform.PerfMainName)

	params.Add("seatGradeNo", block.SeatGradeNo)
	params.Add("seatGradeName", "")
	params.Add("blockId", block.BlockID)
	params.Add("sntv", block.Sntv)
	params.Add("blockTypeCode", "")
	params.Add("floorNo", block.FloorNo)
	params.Add("floorName", block.FloorName)

	params.Add("areaNo", block.AreaNo)

	params.Add("areaName", block.AreaName)
	params.Add("prodTypeCode", i.ProdInform.ProdTypeCode)
	params.Add("flplanTypeCode", "DR0002")
	params.Add("scheduleTypeCode", i.ProdInform.ScheduleTypeCode)
	params.Add("seatTypeCode", "SE0001")
	params.Add("jType", "I")
	params.Add("cardGroupId", "")
	params.Add("cardBpId", "")
	params.Add("cardMid", "")

	params.Add("rsrvStep", "SAT")
	params.Add("zamEnabled", "0")
	params.Add("zamKey", "")
	params.Add("trafficCtrlYn", "N")
	params.Add("netfunnel_key", "")
	params.Add("stvn_view_list", block.StvnViewList)
	params.Add("mapClickYn", "Y")
	params.Add("seatId", seatId)
	params.Add("clipSeatId", "")
	strMap := make(map[string]string, len(params))
	for key, values := range params {
		if len(values) > 0 {
			strMap[key] = values[0]
		}
	}
	//var body = strings.NewReader(`langCd=EN&prodId=209540&pocCode=SC0002&perfTypeCode=GN0001&perfDate=20240414&scheduleNo=100001&sellTypeCode=ST0001&sellCondNo=&perfMainName=2024+EXO+FAN+MEETING%EF%BC%9AONE&seatGradeNo=&seatGradeName=&blockId=188&sntv=%2C%EF%BC%BBA%EF%BC%8EB+GATE%EF%BC%BD402&blockTypeCode=&floorNo=&floorName=&areaNo=%EF%BC%BBA%EF%BC%8EB+GATE%EF%BC%BD402&areaName=%EA%B5%AC%EC%97%AD&prodTypeCode=PT0001&flplanTypeCode=DR0002&scheduleTypeCode=SG0001&seatTypeCode=SE0001&jType=I&cardGroupId=&cardBpId=&cardMid=&rsrvStep=SAT&zamEnabled=0&zamKey=&trafficCtrlYn=N&netfunnel_key=&stvn_view_list=%2C%EF%BC%BBA+GATE%EF%BC%BD401%3B%2C%EF%BC%BBA%EF%BC%8EB+GATE%EF%BC%BD302%3B%2C%EF%BC%BBA%EF%BC%8EB+GATE%EF%BC%BD402%3B%2C%EF%BC%BBB+GATE%EF%BC%BD403%3B%2C%EF%BC%BBB%EF%BC%8EC+GATE%EF%BC%BD303%3B%2C%EF%BC%BBC+GATE%EF%BC%BD404%3B%2C%EF%BC%BBC%EF%BC%8ED+GATE%EF%BC%BD304%3B%2C%EF%BC%BBE+GATE%EF%BC%BD205%3B%2C%EF%BC%BBE+GATE%EF%BC%BD206%3B%2C%EF%BC%BBE+GATE%EF%BC%BD305%3B%2C%EF%BC%BBE+GATE%EF%BC%BD306%3B%2C%EF%BC%BBE+GATE%EF%BC%BDF4%3B%2C%EF%BC%BBE+GATE%EF%BC%BDF5%3B%2C%EF%BC%BBE%EF%BC%8EF+GATE%EF%BC%BD207%3B%2C%EF%BC%BBE%EF%BC%8EF+GATE%EF%BC%BD307%3B%2C%EF%BC%BBF+GATE%EF%BC%BD208%3B%2C%EF%BC%BBF+GATE%EF%BC%BD209%3B%2C%EF%BC%BBF+GATE%EF%BC%BD308%3B%2C%EF%BC%BBF%EF%BC%8EG+GATE%EF%BC%BD309%3B%2C%EF%BC%BBG%EF%BC%8EH+GATE%EF%BC%BD310%3B%2C%EF%BC%BBH+GATE%EF%BC%BD210%3B%2C%EF%BC%BBH%EF%BC%8EJ+GATE%EF%BC%BD311%3B%2C%EF%BC%BBJ+GATE%EF%BC%BD211%3B%2C%EF%BC%BBJ%EF%BC%8EK+GATE%EF%BC%BD212%3B%2C%EF%BC%BBJ%EF%BC%8EK+GATE%EF%BC%BD312%3B%2C%EF%BC%BBJ%EF%BC%8EK+GATE%EF%BC%BD412%3B%2C%EF%BC%BBK+GATE%EF%BC%BD213%3B%2C%EF%BC%BBK%EF%BC%8EL+GATE%EF%BC%BD313%3B%2C%EF%BC%BBK%EF%BC%8EL+GATE%EF%BC%BD413%3B%2C%EF%BC%BBL%EF%BC%8EM+GATE%EF%BC%BD314%3B%2C%EF%BC%BBL%EF%BC%8EM+GATE%EF%BC%BD414%3B%2C%EF%BC%BBM+GATE%EF%BC%BD214%3B%2C%EF%BC%BBM%EF%BC%8EN+GATE%EF%BC%BD315%3B%2C%EF%BC%BBN+GATE%EF%BC%BD215%3B%2C%EF%BC%BBN+GATE%EF%BC%BD415%3B%2C%EF%BC%BBN+GATE%EF%BC%BDF3%3B%2C%EF%BC%BBN%EF%BC%8EP+GATE%EF%BC%BD216%3B%2C%EF%BC%BBN%EF%BC%8EP+GATE%EF%BC%BD316%3B%2C%EF%BC%BBN%EF%BC%8EP+GATE%EF%BC%BD416%3B%2C%EF%BC%BBN%EF%BC%8EP+GATE%EF%BC%BDF1%3B%2C%EF%BC%BBN%EF%BC%8EP+GATE%EF%BC%BDF2%3B%2C%EF%BC%BBP%EF%BC%8EQ+GATE%EF%BC%BD217%3B%2C%EF%BC%BBP%EF%BC%8EQ+GATE%EF%BC%BD317%3B%2C%EF%BC%BBP%EF%BC%8EQ+GATE%EF%BC%BD417%3B%2C%EF%BC%BBQ%EF%BC%8ER+GATE%EF%BC%BD218%3B%2C%EF%BC%BBQ%EF%BC%8ER+GATE%EF%BC%BD318%3B%2C%EF%BC%BBQ%EF%BC%8ER+GATE%EF%BC%BD418%3B%2C%EF%BC%BBR+GATE%EF%BC%BD419%3B%2C%EF%BC%BBS%EF%BC%8ET+GATE%EF%BC%BD320%3B%2C%EF%BC%BBS%EF%BC%8ET+GATE%EF%BC%BD420%3B%2C%EF%BC%BBT%EF%BC%8EA+GATE%EF%BC%BD301&mapClickYn=Y&seatId=` + seatId)
	var body = strings.NewReader(`langCd=EN&prodId=` + strconv.Itoa(i.ProdInform.ProdID) + `&pocCode=` + i.ProdInform.PocCode + `&perfTypeCode=` + i.ProdInform.ProdTypeCode + `&perfDate=` + i.ProdInform.PerfStartDay + `&scheduleNo=` + strconv.Itoa(i.ScheduleNo) + `
				&sellTypeCode=` + i.SchList[0].SellTypeCode + `&sellCondNo=&perfMainName=` + i.ProdInform.PerfMainName + `&seatGradeNo=` + block.SeatGradeNo + `&seatGradeName=&blockId=` + block.BlockID + `&sntv=` + block.Sntv + `&blockTypeCode=&floorNo=` + block.FloorNo + `&floorName=` + block.FloorName + `
				&areaNo=` + block.AreaNo + `&areaName=` + block.AreaName + `&prodTypeCode=` + i.ProdInform.ProdTypeCode + `&flplanTypeCode=DR0002&scheduleTypeCode=` + i.ProdInform.ScheduleTypeCode + `&seatTypeCode=SE0001&jType=I&cardGroupId=&cardBpId=&cardMid=&rsrvStep=SAT&zamEnabled=0&zamKey=&trafficCtrlYn=N&netfunnel_key=&stvn_view_list=` + url.QueryEscape(block.StvnViewList) + `&mapClickYn=Y&seatId=` + seatId)
	fmt.Println(body)
	_ = strings.NewReader(params.Encode())

	r, _ := httpGet.HttpMelonQueryPost(c, "https://tkglobal.melon.com/tktapi/glb/product/tickettype.json?v=1", strMap)
	if string(r) == `{"code":"9000","staticDomain":null,"message":"로그인 상태가 아닙니다.","httpsDomain":null,"httpDomain":null}` {
		return saveDataReq, fmt.Errorf("cookie过期")
	}

	var t model.TypeTicketResp

	err = json.Unmarshal(r, &t)
	if err != nil {
		fmt.Println(err)
	} else {
		if len(t.SeatGradeList) > 0 {
			//fmt.Println(string(r))
			logging.Info(t)
			for _, v := range t.SeatGradeList {
				for _, b := range v.ProdTicketTypeList {

					if strconv.Itoa(b.SeatGradeNo) == block.SeatGradeNo {
						saveDataReq.PriceNo = strconv.Itoa(b.PriceNo)
					}
				}

			}

			saveDataReq.SeatGradeName = block.SeatGradeName
			saveDataReq.PayAmt = t.SeatGradeList[0].ProdTicketTypeList[0].TicketTypePrice + 2000
			saveDataReq.PriceName = t.SeatGradeList[0].ProdTicketTypeList[0].KrPriceName
			fmt.Println(t.SeatGradeList[0].ProdTicketTypeList[0].TicketTypePrice)
			saveDataReq.BasePrice = strconv.Itoa(t.SeatGradeList[0].ProdTicketTypeList[0].TicketTypePrice)
		} else {
			fmt.Println(string(r))
		}

	}

	return saveDataReq, err
}
func GetSeatList(c *req.Client, r model.SeatReq) (seat model.SeatList, err error) {
	//logging.Info("获取座位列表：seatMapList")
	if r.BlockID == "0" {
		r.BlockID = ""
	}
	//fmt.Println("--")

	fmt.Println(`prodId=` + r.ProdID + `&scheduleNo=` + r.ScheduleNo + `&blockId=` + r.BlockID + `&pocCode=SC0002&corpCodeNo=`)

	u := "https://tkglobal.melon.com/tktapi/product/seat/seatMapList.json?v=1&callback=getSeatListCallBack"
	//respByte, err := http2MelonQuery(u, params)
	respByte, err := httpGet.HttpMelonQueryPost(c, u, map[string]string{
		"prodId":     r.ProdID,
		"scheduleNo": r.ScheduleNo,
		"blockId":    r.BlockID,
		"pocCode":    "SC0002",
		"corpCodeNo": "",
	})

	if err != nil {

		return
	}

	if strings.Contains(string(respByte), "403 Forbidden") {
		fmt.Println(string(respByte))
		logging.Error("403")
		return seat, fmt.Errorf("403")
	}
	if !strings.Contains(string(respByte), "/**/getSeatListCallBack(") {
		logging.Error(string(respByte))
		fmt.Println(string(respByte))
	}
	rs := strings.Trim(string(respByte), "/**/getSeatListCallBack(")
	re := strings.TrimRight(rs, ");")
	err = json.Unmarshal([]byte(re), &seat)
	if err != nil {
		logging.Error(err)
		//logging.Error(string(respByte))
	}

	return
}
func stepTicket(c *req.Client, prodId string, scheduleNo string, data model.ProdLimitData) {
	params := url.Values{}
	params.Add("prodId", prodId)
	params.Add("scheduleNo", scheduleNo)
	params.Add("flplanTypeCode", "DR0002")
	params.Add("seatTypeCode", "SE0001")
	params.Add("encryptedSeatIds", data.EncryptedSeatIds)
	params.Add("interlockTypeCode", "")
	params.Add("interlockTid", "0")
	params.Add("seatIds", data.SeatIds)
	//body := strings.NewReader(params.Encode())
	strMap := make(map[string]string, len(params))
	for key, values := range params {
		if len(values) > 0 {
			strMap[key] = values[0]
		}
	}
	_, err := httpGet.HttpMelonQueryPost(c, "https://tkglobal.melon.com/reservation/popup/stepTicket.htm", strMap)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Printf(string(respByte))
}
