package summary

import (
	"encoding/json"
	"fmt"
	"github.com/imroc/req/v3"
	"melonNew/global"
	"melonNew/httpGet"
	"melonNew/logging"
	"melonNew/model"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func InformProdSch(c *req.Client, prodId string, pocCode string, scheduleNo string, sellTypeCode string) (i model.PromReqData) {
	logging.Info("获取座位信息：informProdSch")
	params := url.Values{}
	params.Add("prodId", prodId)
	params.Add("pocCode", pocCode)
	params.Add("scheduleNo", scheduleNo)
	params.Add("sellTypeCode", sellTypeCode)
	params.Add("seatGradeNo", "ST0001")

	respByte, _ := httpGet.HttpMelonQueryPost(c, "https://tkglobal.melon.com/tktapi/product/informProdSch.json?v=1", map[string]string{
		"prodId":       prodId,
		"pocCode":      pocCode,
		"scheduleNo":   scheduleNo,
		"sellTypeCode": sellTypeCode,
		"seatGradeNo":  "ST0001",
	})
	//respByte, _ := http2MelonQuery("https://tkglobal.melon.com/tktapi/product/informProdSch.json?v=1", params)
	json.Unmarshal(respByte, &i)
	fmt.Println(i.ProdInform.PerfStartDay)
	return
}
func S1(c *req.Client, prodId string, pocCode string, scheduleNo string, startDay string) (model.SummaryResp, model.AreaMaps) {
	var resp []byte
	var rs []byte
	dir, _ := os.Getwd()
	if !fileExists(dir + "/" + prodId + scheduleNo + ".txt") {
		resp = summary1(c, prodId, pocCode, scheduleNo, startDay)
	} else {
		resp, _ = os.ReadFile(dir + "/" + prodId + scheduleNo + ".txt")

	}
	if !fileExists(dir + "/" + prodId + scheduleNo + "Area.txt") {
		rs = areaMapF(c, prodId, pocCode, scheduleNo)
	} else {
		rs, _ = os.ReadFile(dir + "/" + prodId + scheduleNo + "Area.txt")

	}
	//aMap := areaMap(prodId, pocCode, scheduleNo)
	var aMap model.AreaMaps
	var d1 model.SummaryResp

	err := json.Unmarshal(rs, &aMap)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(resp, &d1)
	if err != nil {
		fmt.Println(err)
	}
	return d1, aMap
}
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
func Summary(c *req.Client, prodId string, pocCode string, scheduleNo string, startDay string, d1 model.SummaryResp, aMap model.AreaMaps) (r []model.SeatReq, err error) {
	var seat model.SeatReq

	var all []model.SeatReq

	//var resp []byte
	//var rs []byte
	logging.Info("查询全图票情况：summary")

	fmt.Println("执行Summary")
	//for _, v := range d1.Summary {
	//	//fmt.Println(v.SeatGradeName)
	//}
	params := url.Values{}

	params.Add("prodId", prodId)
	params.Add("pocCode", pocCode)
	params.Add("scheduleNo", scheduleNo)
	//params.Add("seatGradeNo", "11031")

	//respByte, err := http2MelonQuery("https://tkglobal.melon.com/tktapi/product/block/summary.json?v=1", params)

	respByte, err := httpGet.HttpMelonQueryPost(c, "https://tkglobal.melon.com/tktapi/product/block/summary.json?v=1", map[string]string{
		"prodId":     prodId,
		"pocCode":    pocCode,
		"scheduleNo": scheduleNo,
		"langCd":     "EN",
	})
	if err != nil {
		return r, err
	}

	//fmt.Println(string(respByte))

	//respByte, err := os.ReadFile(prodId + "2.txt")
	var d model.JSONData
	fmt.Println(startDay)
	//fmt.Println(string(respByte))
	err = json.Unmarshal(respByte, &d)
	//fmt.Println(string(respByte))
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(d.Summary)
	fmt.Println(len(d.Summary))

	for k, v := range d.Summary {

		for a, b := range d1.Summary {

			if b.SeatGradeNo == v.SeatGradeNo {
				seat.StvnViewList = d1.Summary[a].SntvList
				seat.BasePrice = strconv.Itoa(b.BasePrice)

			}
		}

		seat.ScheduleNo = scheduleNo

		//seat.BlockID = strconv.Itoa(d1.Summary[k].BlockID)
		seat.SeatGradeNo = strconv.Itoa(d.Summary[k].SeatGradeNo)
		seat.ProdID = prodId
		seat.AreaNo = d.Summary[k].AreaNo
		seat.Sntv = d.Summary[k].Sntv
		seat.FloorNo = v.FloorNo
		seat.FloorName = v.FloorName
		seat.AreaName = v.AreaName
		seat.SeatGradeName = v.SeatGradeName

		seat.PerfStartDay = startDay
		seat.BlockID = ""
		for _, n := range aMap.SeatData.Da.Sb {
			if len(seat.BlockID) > 0 {

				if n.Sntv.A == v.AreaNo && n.Sntv.F == v.FloorNo {
					seat.BlockID = seat.BlockID + "," + strconv.Itoa(n.Sbid)
				}
			} else {
				if n.Sntv.A == v.AreaNo && n.Sntv.F == v.FloorNo {
					seat.BlockID = strconv.Itoa(n.Sbid)
				}
			}

		}
		AreaNos := global.GlobalConfig.GetString("areaNos")
		if len(AreaNos) > 0 {
			for _, area := range strings.Split(AreaNos, ",") {
				fmt.Println(seat.AreaNo, seat.BlockID)
				if strings.Contains(seat.AreaNo, area) {
					fmt.Println(seat)
					all = append(all, seat)
				}
				//if strings.Contains(seat.AreaNo, area) {
				//	fmt.Println(seat)
				//	all = append(all, seat)
				//}
			}
		}

		fmt.Println(len(all))

	}

	return all, nil

}
func areaMapF(c *req.Client, prodId string, pocCode string, scheduleNo string) []byte {

	u := "https://tkglobal.melon.com/tktapi/glb/product/getAreaMap.json?v=1"

	resp, err := httpGet.HttpMelonQueryPost(c, u, map[string]string{
		"prodId":     prodId,
		"pocCode":    "SC0002",
		"scheduleNo": scheduleNo,
	})

	if err != nil {
		fmt.Println(err)

	}

	dir, _ := os.Getwd()
	openFile, e := os.OpenFile(dir+"/"+prodId+scheduleNo+"Area.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if e != nil {
		fmt.Println(e)
	}
	defer openFile.Close()

	_, _ = openFile.Write(resp)

	//fmt.Println(r.SeatData.Da.Sb)
	//for _, v := range r.SeatData.Da.Sb {
	//	fmt.Println(v.Sntv)
	//}
	return resp
}
func summary1(c *req.Client, prodId string, pocCode string, scheduleNo string, startDay string) []byte {
	param := url.Values{}

	param.Add("prodId", prodId)
	param.Add("pocCode", pocCode)
	param.Add("scheduleNo", scheduleNo)
	param.Add("perfDate", startDay)
	param.Add("langCd", "EN")

	resp, _ := httpGet.HttpMelonQueryPost(c, "https://tkglobal.melon.com/tktapi/glb/product/summary.json?v=1", map[string]string{
		"prodId":     prodId,
		"pocCode":    pocCode,
		"scheduleNo": scheduleNo,
		"perfDate":   startDay,
		"langCd":     "EN",
	})
	dir, _ := os.Getwd()
	openFile, e := os.OpenFile(dir+"/"+prodId+scheduleNo+".txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if e != nil {
		fmt.Println(e)
	}
	defer openFile.Close()

	_, _ = openFile.Write(resp)
	return resp
}
