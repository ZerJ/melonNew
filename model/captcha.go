package model

type ProdKey struct {
	Code          string      `json:"code"`
	NflActID      string      `json:"nflActId"`
	StaticDomain  interface{} `json:"staticDomain"`
	HTTPSDomain   interface{} `json:"httpsDomain"`
	TrafficCtrlYn string      `json:"trafficCtrlYn"`
	Key           string      `json:"key"`
	HTTPDomain    interface{} `json:"httpDomain"`
}
type CheckCaptcha struct {
	UserCaptStr  string `json:"userCaptStr"`
	Chkcapt      string `json:"chkcapt"`
	ProdId       string `json:"prodId"`
	ScheduleNo   string `json:"scheduleNo"`
	PocCode      string `json:"pocCode"`
	SellTypeCode string `json:"sellTypeCode"`
}
type Code struct {
	Code string `json:"code"`
}
type CaptchaData struct {
	Code         string `json:"code"`
	StaticDomain string `json:"staticDomain"`
	HttpsDomain  string `json:"httpsDomain"`
	Data         string `json:"data"`
	HttpDomain   string `json:"httpsomain"`
}
