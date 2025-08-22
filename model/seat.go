package model

type PromReqData struct {
	Code         string      `json:"code"`
	StaticDomain interface{} `json:"staticDomain"`
	ScheduleNo   int         `json:"scheduleNo"`
	HTTPSDomain  interface{} `json:"httpsDomain"`
	IsMobile     bool        `json:"isMobile"`
	ProdInform   struct {
		ProdID            int         `json:"prodId"`
		PerfMainName      string      `json:"perfMainName"`
		PlaceID           int         `json:"placeId"`
		PlaceCodeName     string      `json:"placeCodeName"`
		HallID            int         `json:"hallId"`
		HallCodeName      string      `json:"hallCodeName"`
		BpNo              int         `json:"bpNo"`
		ProdTypeCode      string      `json:"prodTypeCode"`
		GenrlRsrvYn       string      `json:"genrlRsrvYn"`
		BefRsrvYn         string      `json:"befRsrvYn"`
		PickRsrvYn        string      `json:"pickRsrvYn"`
		WishRsrvYn        string      `json:"wishRsrvYn"`
		ScheduleTypeCode  string      `json:"scheduleTypeCode"`
		ScheduleDisplayYn interface{} `json:"scheduleDisplayYn"`
		SellStateCode     string      `json:"sellStateCode"`
		CancelRsrvYn      string      `json:"cancelRsrvYn"`
		SeatMdfYn         string      `json:"seatMdfYn"`
		FSeatLockYn       string      `json:"fSeatLockYn"`
		RsrvPaySepYn      string      `json:"rsrvPaySepYn"`
		VbankCloseDt      string      `json:"vbankCloseDt"`
		VbankCloseTime    string      `json:"vbankCloseTime"`
		VbankYn           string      `json:"vbankYn"`
		MoblPaymtYn       string      `json:"moblPaymtYn"`
		SpotOperYn        string      `json:"spotOperYn"`
		VolumeTypeCode    string      `json:"volumeTypeCode"`
		LimitVolume       int         `json:"limitVolume"`
		SalesUserID       string      `json:"salesUserId"`
		PostDelvyYn       string      `json:"postDelvyYn"`
		SpotRecvYn        string      `json:"spotRecvYn"`
		MobileRecvYn      string      `json:"mobileRecvYn"`
		DelvyFee          int         `json:"delvyFee"`
		DelvyGuide        string      `json:"delvyGuide"`
		BatchDelvyDt      string      `json:"batchDelvyDt"`
		DelvyClosePolicy  string      `json:"delvyClosePolicy"`
		TrafficCtrlYn     string      `json:"trafficCtrlYn"`
		AddInfoYn         string      `json:"addInfoYn"`
		ParcelRecvYn      string      `json:"parcelRecvYn"`
		GlobalDelvyYn     string      `json:"globalDelvyYn"`
		WeightCode        string      `json:"weightCode"`
		PayKakaoYn        string      `json:"payKakaoYn"`
		IsMdProduct       string      `json:"isMdProduct"`
		SeatCntDisplayYn  string      `json:"seatCntDisplayYn"`
		PerfTypeCode      string      `json:"perfTypeCode"`
		UseYn             string      `json:"useYn"`
		PocCode           string      `json:"pocCode"`
		RsrvFee           int         `json:"rsrvFee"`
		MaxCancelFeeVal   int         `json:"maxCancelFeeVal"`
		ScheduleNo        int         `json:"scheduleNo"`
		PerfStartDay      string      `json:"perfStartDay"`
		PerfStartTime     string      `json:"perfStartTime"`
		FlplanID          string      `json:"flplanId"`
		DelvyCloseDt      string      `json:"delvyCloseDt"`
		CancelCloseDt     string      `json:"cancelCloseDt"`
		RsrvStartDt       string      `json:"rsrvStartDt"`
		RsrvEndDt         string      `json:"rsrvEndDt"`
		BizTypeFlg        string      `json:"bizTypeFlg"`
		FlplanTypeCode    string      `json:"flplanTypeCode"`
		RsrvAvailYn       string      `json:"rsrvAvailYn"`
		MelonVipYn        string      `json:"melonVipYn"`
		IsVipRsrvFee      interface{} `json:"isVipRsrvFee"`
		DeductionYn       string      `json:"deductionYn"`
	} `json:"prodInform"`
	SchList []struct {
		ProdID           int         `json:"prodId"`
		ScheduleNo       int         `json:"scheduleNo"`
		SellFlplanID     int         `json:"sellFlplanId"`
		OrderNo          int         `json:"orderNo"`
		PerfStartDay     string      `json:"perfStartDay"`
		PerfStartTime    string      `json:"perfStartTime"`
		DelvyCloseDt     interface{} `json:"delvyCloseDt"`
		SpotDelvyCloseDt interface{} `json:"spotDelvyCloseDt"`
		CancelCloseDt    interface{} `json:"cancelCloseDt"`
		UseYn            interface{} `json:"useYn"`
		TotSeatCntlk     int         `json:"totSeatCntlk"`
		SellSeatCntlk    int         `json:"sellSeatCntlk"`
		RendrSeatCntlk   int         `json:"rendrSeatCntlk"`
		LockSeatCntlk    int         `json:"lockSeatCntlk"`
		RealSeatCntlk    int         `json:"realSeatCntlk"`
		RsrvAvailYn      interface{} `json:"rsrvAvailYn"`
		FlplanTypeCode   interface{} `json:"flplanTypeCode"`
		SellStateCode    interface{} `json:"sellStateCode"`
		SellTypeCode     string      `json:"sellTypeCode"`
		TrafficCtrlYn    interface{} `json:"trafficCtrlYn"`
		PostDelvyYn      interface{} `json:"postDelvyYn"`
		SpotRecvYn       interface{} `json:"spotRecvYn"`
		MobileRecvYn     interface{} `json:"mobileRecvYn"`
		VbankYn          interface{} `json:"vbankYn"`
		MoblPaymtYn      interface{} `json:"moblPaymtYn"`
		PayCardYn        interface{} `json:"payCardYn"`
		PayMobileYn      interface{} `json:"payMobileYn"`
		PayBankYn        interface{} `json:"payBankYn"`
		PayKakaoYn       interface{} `json:"payKakaoYn"`
		CardPointUseYn   interface{} `json:"cardPointUseYn"`
		ClipYn           interface{} `json:"clipYn"`
		SeatCntVOList    interface{} `json:"seatCntVOList"`
		SeatGradelistVO  interface{} `json:"seatGradelistVO"`
		PocCode          string      `json:"pocCode"`
		HasPresaleAdvtk  string      `json:"hasPresaleAdvtk"`
		RsrvStartDt      string      `json:"rsrvStartDt"`
		RsrvEndDt        interface{} `json:"rsrvEndDt"`
		GroupSch         interface{} `json:"groupSch"`
		SeatCntDisplayYn string      `json:"seatCntDisplayYn"`
		Casting          interface{} `json:"casting"`
	} `json:"schList"`
	HTTPDomain interface{} `json:"httpDomain"`
}
type SummaryResp struct {
	Code              string          `json:"code"`
	HttpDomain        string          `json:"httpDomain"`
	HttpsDomain       string          `json:"httpsDomain"`
	InterlockTypeCode string          `json:"interlockTypeCode"`
	StaticDomain      string          `json:"staticDomain"`
	Summary           []AutoGenerated `json:"summary"`
}
type AutoGenerated struct {
	ProdID              int         `json:"prodId"`
	ScheduleNo          int         `json:"scheduleNo"`
	BlockID             int         `json:"blockId"`
	SeatGradeNo         int         `json:"seatGradeNo"`
	TotSeatCntlk        int         `json:"totSeatCntlk"`
	SellSeatCntlk       int         `json:"sellSeatCntlk"`
	RendrSeatCntlk      int         `json:"rendrSeatCntlk"`
	LockSeatCntlk       int         `json:"lockSeatCntlk"`
	RealSeatCntlk       int         `json:"realSeatCntlk"`
	RegUserID           interface{} `json:"regUserId"`
	RegUserName         interface{} `json:"regUserName"`
	RegDate             interface{} `json:"regDate"`
	MdfUserID           interface{} `json:"mdfUserId"`
	MdfUserName         interface{} `json:"mdfUserName"`
	MdfDate             interface{} `json:"mdfDate"`
	SeatGradeName       string      `json:"seatGradeName"`
	GradeColorVal       string      `json:"gradeColorVal"`
	FloorNo             string      `json:"floorNo"`
	FloorName           string      `json:"floorName"`
	AreaNo              string      `json:"areaNo"`
	AreaName            string      `json:"areaName"`
	SntvList            string      `json:"sntvList"`
	Sntv                string      `json:"sntv"`
	BlockTypeCode       string      `json:"blockTypeCode"`
	PerfStartDay        string      `json:"perfStartDay"`
	PerfStartTime       string      `json:"perfStartTime"`
	BasePrice           int         `json:"basePrice"`
	SejongSeatGradeCode interface{} `json:"sejongSeatGradeCode"`
}
type AreaMaps struct {
	SeatData struct {
		St []struct {
			Ss   []interface{} `json:"ss"`
			Sbid int           `json:"sbid"`
		} `json:"st"`
		Plid string `json:"plid"`
		Bu   string `json:"bu"`
		Plhl string `json:"plhl"`
		Im   string `json:"im"`
		Ms   struct {
			Width  string `json:"width"`
			Height string `json:"height"`
		} `json:"ms"`
		Mt  string `json:"mt"`
		Snt struct {
			A struct {
				Use  string `json:"use"`
				Name string `json:"name"`
			} `json:"a"`
			R struct {
				Use  string `json:"use"`
				Name string `json:"name"`
			} `json:"r"`
			E struct {
				Use  string `json:"use"`
				Name string `json:"name"`
			} `json:"e"`
			F struct {
				Use  string `json:"use"`
				Name string `json:"name"`
			} `json:"f"`
			N struct {
				Use  string `json:"use"`
				Name string `json:"name"`
			} `json:"n"`
		} `json:"snt"`
		Pid string `json:"pid"`
		Da  struct {
			Bb []interface{} `json:"bb"`
			Zb []struct {
				Dt  string      `json:"dt"`
				Cd  interface{} `json:"cd"`
				R   interface{} `json:"r"`
				Zid int         `json:"zid"`
				Ot  string      `json:"ot"`
				Lc  string      `json:"lc,omitempty"`
				Ls  string      `json:"ls,omitempty"`
				Snt struct {
					A string `json:"a"`
					R string `json:"r"`
					E string `json:"e"`
					F string `json:"f"`
				} `json:"snt,omitempty"`
				Fc         string `json:"fc,omitempty"`
				FontColor  string `json:"font_color,omitempty"`
				FontSize   string `json:"font_size,omitempty"`
				Text       string `json:"text,omitempty"`
				FontName   string `json:"font_name,omitempty"`
				FontEffect string `json:"font_effect,omitempty"`
			} `json:"zb"`
			Sb []struct {
				Cc int `json:"cc"`
				Cd struct {
					W float64 `json:"w"`
					X float64 `json:"x"`
					H float64 `json:"h"`
					Y float64 `json:"y"`
				} `json:"cd"`
				CSS  interface{} `json:"css"`
				Rt   int         `json:"rt"`
				Sbt  string      `json:"sbt"`
				Ls   string      `json:"ls"`
				Sntv struct {
					A string `json:"a"`
					R string `json:"r"`
					E string `json:"e"`
					F string `json:"f"`
				} `json:"sntv,omitempty"`
				It   string      `json:"it"`
				Iv   interface{} `json:"iv"`
				Ssn  int         `json:"ssn"`
				Rc   int         `json:"rc"`
				Rss  interface{} `json:"rss"`
				Sst  string      `json:"sst,omitempty"`
				Spt  string      `json:"spt"`
				Snt  string      `json:"snt"`
				If   string      `json:"if"`
				Sbid int         `json:"sbid"`
			} `json:"sb"`
		} `json:"da"`
		Pnm  string `json:"pnm"`
		Plnm string `json:"plnm"`
	} `json:"seatData"`
	Code         string      `json:"code"`
	StaticDomain interface{} `json:"staticDomain"`
	HTTPSDomain  interface{} `json:"httpsDomain"`
	HTTPDomain   interface{} `json:"httpDomain"`
}
type SeatReq struct {
	ProdID        string `json:"prodId"`
	ScheduleNo    string `json:"scheduleNo"`
	BlockID       string `json:"blockId"`
	SeatGradeNo   string `json:"seatGradeNo"`
	SeatGradeName string `json:"seatGradeName"`
	Sntv          string `json:"sntv"`
	AreaNo        string `json:"areaNo"`
	BlockTypeCode string `json:"blockTypeCode"`
	AreaName      string `json:"areaName"`
	StvnViewList  string `json:"stvnViewList"`
	FloorNo       string `json:"floorNo"`
	FloorName     string `json:"floorName"`
	BasePrice     string `json:"basePrice"`
	PerfStartDay  string `json:"perfStartDay"`
}
type JSONData struct {
	Summary []struct {
		FloorNo       string `json:"floorNo"`
		FloorName     string `json:"floorName"`
		AreaNo        string `json:"areaNo"`
		AreaName      string `json:"areaName"`
		Sntv          string `json:"sntv"`
		SeatGradeNo   int    `json:"seatGradeNo"`
		SeatGradeName string `json:"seatGradeName"`
		RealSeatCntlk int    `json:"realSeatCntlk"`
	} `json:"summary"`
	InterlockTypeCode string      `json:"interlockTypeCode"`
	Code              string      `json:"code"`
	StaticDomain      interface{} `json:"staticDomain"`
	HTTPSDomain       interface{} `json:"httpsDomain"`
	HTTPDomain        interface{} `json:"httpDomain"`
}
type SeatList struct {
	HttpDomain   string  `json:"httpDomain"`
	HttpsDomain  string  `json:"httpsDomain"`
	Params       Param   `json:"params"`
	SeatData     Seat    `json:"seatData"`
	SeatInxData  InxData `json:"seatInxData"`
	StaticDomain string  `json:"staticDomain"`
}
type Param struct {
	V             string   `json:"V"`
	PRODID        string   `json:"PRODID"`
	SCHEDULENO    string   `json:"SCHEDULENO"`
	BLOCKID       string   `json:"BLOCKID"`
	CALLBACK      string   `json:"CALLBACK"`
	POCCODE       string   `json:"POCCODE"`
	SNTV          string   `json:"SNTV"`
	BLOCKTYPECODE string   `json:"BLOCKTYPECODE"`
	AREANO        string   `json:"AREANO"`
	SEATGRADENO   string   `json:"SEATGRADENO"`
	BLOCKIDS      []string `json:"BLOCKIDS"`
}
type Seat struct {
	St []struct {
		Sbid int `json:"sbid"`
		Ss   []struct {
			Sid  string      `json:"sid"`
			Rn   string      `json:"rn"`
			Sn   string      `json:"sn"`
			Sc   int         `json:"sc"`
			Snmm interface{} `json:"snmm"`
			Snm  string      `json:"snm"`
			Cd   struct {
				X float64 `json:"x"`
				Y float64 `json:"y"`
			} `json:"cd"`
			Uf      string      `json:"uf"`
			Sts     string      `json:"sts"`
			Gd      string      `json:"gd"`
			Gc      string      `json:"gc"`
			Cl      interface{} `json:"cl"`
			Std     string      `json:"std"`
			Dc      string      `json:"dc"`
			Slc     int         `json:"slc"`
			Sl      string      `json:"sl"`
			Csid    interface{} `json:"csid"`
			Csnm    interface{} `json:"csnm"`
			Gn      string      `json:"gn"`
			Ntc     int         `json:"ntc"`
			Amslc   int         `json:"amslc"`
			Aslc    int         `json:"aslc"`
			Itc     interface{} `json:"itc"`
			Stn     interface{} `json:"stn"`
			NauseYn interface{} `json:"nauseYn"`
		} `json:"ss"`
	} `json:"st"`
	Plid string `json:"plid"`
	Bu   string `json:"bu"`
	Plhl string `json:"plhl"`
	Im   string `json:"im"`
	Ms   struct {
		Width  string `json:"width"`
		Height string `json:"height"`
	} `json:"ms"`
	Mt  string `json:"mt"`
	Snt struct {
		A struct {
			Use  string `json:"use"`
			Name string `json:"name"`
		} `json:"a"`
		R struct {
			Use  string `json:"use"`
			Name string `json:"name"`
		} `json:"r"`
		E struct {
			Use  string `json:"use"`
			Name string `json:"name"`
		} `json:"e"`
		F struct {
			Use  string `json:"use"`
			Name string `json:"name"`
		} `json:"f"`
		N struct {
			Use  string `json:"use"`
			Name string `json:"name"`
		} `json:"n"`
	} `json:"snt"`
	Pid string `json:"pid"`
	Da  struct {
		Bb []interface{} `json:"bb"`
		Zb []interface{} `json:"zb"`
		Sb []struct {
			Cc int `json:"cc"`
			Cd struct {
				W float64 `json:"w"`
				X float64 `json:"x"`
				H float64 `json:"h"`
				Y float64 `json:"y"`
			} `json:"cd"`
			CSS  float64 `json:"css"`
			Rt   float64 `json:"rt"`
			Sbt  string  `json:"sbt"`
			Ls   string  `json:"ls"`
			Sntv struct {
				A string `json:"a"`
				R string `json:"r"`
				E string `json:"e"`
				F string `json:"f"`
			} `json:"sntv"`
			It   string  `json:"it"`
			Iv   string  `json:"iv"`
			Ssn  float64 `json:"ssn"`
			Rc   float64 `json:"rc"`
			Rss  float64 `json:"rss"`
			Sst  string  `json:"sst"`
			Spt  string  `json:"spt"`
			Snt  string  `json:"snt"`
			If   string  `json:"if"`
			Sbid int     `json:"sbid"`
		} `json:"sb"`
	} `json:"da"`
	Pnm  string `json:"pnm"`
	Plnm string `json:"plnm"`
}

type InxData struct {
	St []struct {
		Ss   []interface{} `json:"ss"`
		Sbid int           `json:"sbid"`
	} `json:"st"`
	Plid string `json:"plid"`
	Bu   string `json:"bu"`
	Plhl string `json:"plhl"`
	Im   string `json:"im"`
	Ms   struct {
		Width  string `json:"width"`
		Height string `json:"height"`
	} `json:"ms"`
	Mt  string `json:"mt"`
	Snt struct {
		A struct {
			Use  string `json:"use"`
			Name string `json:"name"`
		} `json:"a"`
		R struct {
			Use  string `json:"use"`
			Name string `json:"name"`
		} `json:"r"`
		E struct {
			Use  string `json:"use"`
			Name string `json:"name"`
		} `json:"e"`
		F struct {
			Use  string `json:"use"`
			Name string `json:"name"`
		} `json:"f"`
		N struct {
			Use  string `json:"use"`
			Name string `json:"name"`
		} `json:"n"`
	} `json:"snt"`
	Pid string `json:"pid"`
	Da  struct {
		Bb []interface{} `json:"bb"`
		Zb []struct {
			Dt  string `json:"dt"`
			Cd  string `json:"cd"`
			R   int    `json:"r"`
			Zid int    `json:"zid"`
			Ot  string `json:"ot"`
			Lc  string `json:"lc,omitempty"`
			Ls  string `json:"ls,omitempty"`
			Snt struct {
				A string `json:"a"`
				R string `json:"r"`
				E string `json:"e"`
				F string `json:"f"`
			} `json:"snt,omitempty"`
			Fc         string `json:"fc,omitempty"`
			FontColor  string `json:"font_color,omitempty"`
			FontSize   string `json:"font_size,omitempty"`
			Text       string `json:"text,omitempty"`
			FontName   string `json:"font_name,omitempty"`
			FontEffect string `json:"font_effect,omitempty"`
		} `json:"zb"`
		Sb []struct {
			Cc int `json:"cc"`
			Cd struct {
				W int     `json:"w"`
				X float64 `json:"x"`
				H int     `json:"h"`
				Y float64 `json:"y"`
			} `json:"cd"`
			CSS  int    `json:"css"`
			Rt   int    `json:"rt"`
			Sbt  string `json:"sbt"`
			Ls   string `json:"ls"`
			Sntv struct {
				A string `json:"a"`
				R string `json:"r"`
				E string `json:"e"`
				F string `json:"f"`
			} `json:"sntv"`
			It   string `json:"it"`
			Iv   string `json:"iv"`
			Ssn  int    `json:"ssn"`
			Rc   int    `json:"rc"`
			Rss  int    `json:"rss"`
			Sst  string `json:"sst"`
			Spt  string `json:"spt"`
			Snt  string `json:"snt"`
			If   string `json:"if"`
			Sbid int    `json:"sbid"`
		} `json:"sb"`
	} `json:"da"`
	Pnm  string `json:"pnm"`
	Plnm string `json:"plnm"`
}
type ProdLimitData struct {
	InterlockTypeCode string      `json:"interlockTypeCode"`
	EncryptedSeatIds  string      `json:"encryptedSeatIds"`
	Result            string      `json:"result"`
	StaticDomain      interface{} `json:"staticDomain"`
	HTTPSDomain       interface{} `json:"httpsDomain"`
	InterlockTid      string      `json:"interlockTid"`
	VolumeTypeCode    interface{} `json:"volumeTypeCode"`
	Message           string      `json:"message"`
	HTTPDomain        interface{} `json:"httpDomain"`
	SeatIds           string      `json:"seatIds"`
}
type SaveReqData struct {
	ProdId                    string `json:"prodId"`
	PocCode                   string `json:"pocCode"`
	ScheduleNo                string `json:"PocCode"`
	PayAmt                    int    `json:"payAmt"`
	PriceNo                   string `json:"priceNo"`
	SeatId                    string `json:"seatId"`
	SeatInfoListWithPriceType string `json:"seatInfoListWithPriceType"`
	FirstSeatId               string `json:"firstSeatId"`
	SellTypeCode              string `json:"sellTypeCode"`
	Chkcapt                   string `json:"chkcapt"`
	R                         string `json:"r"`
	S                         string `json:"s"`
	SeatGradeName             string `json:"seatGradeName"`
	Price                     string `json:"price"`
	Sntv                      string `json:"sntv"`
	PriceName                 string `json:"priceName"`
	BasePrice                 string `json:"basePrice"`
	FloorNo                   string `json:"floorNo"`
	AreaNo                    string `json:"areaNo"`
	AreaName                  string `json:"areaName"`
	FloorName                 string `json:"floorName"`
}
type TypeTicketResp struct {
	Code          string        `json:"code"`
	StaticDomain  interface{}   `json:"staticDomain"`
	HTTPSDomain   interface{}   `json:"httpsDomain"`
	AdvtkList     []interface{} `json:"advtkList"`
	SeatGradeList []struct {
		ProdID             int         `json:"prodId"`
		ScheduleNo         string      `json:"scheduleNo"`
		SeatID             interface{} `json:"seatId"`
		FlplanID           int         `json:"flplanId"`
		SeatGradeNo        int         `json:"seatGradeNo"`
		RegUserID          interface{} `json:"regUserId"`
		RegUserName        interface{} `json:"regUserName"`
		RegDate            interface{} `json:"regDate"`
		MdfUserID          interface{} `json:"mdfUserId"`
		MdfUserName        interface{} `json:"mdfUserName"`
		MdfDate            interface{} `json:"mdfDate"`
		BasePrice          int         `json:"basePrice"`
		LockUserID         interface{} `json:"lockUserId"`
		LockStartDt        interface{} `json:"lockStartDt"`
		LockEndDt          interface{} `json:"lockEndDt"`
		TotVisitorCnt      int         `json:"totVisitorCnt"`
		VisitorCnt         int         `json:"visitorCnt"`
		PriceNo            int         `json:"priceNo"`
		DcTypeCode         interface{} `json:"dcTypeCode"`
		DcValue            int         `json:"dcValue"`
		DcPrice            int         `json:"dcPrice"`
		DupDcYn            interface{} `json:"dupDcYn"`
		AdvtkNo            interface{} `json:"advtkNo"`
		RsrvFee            int         `json:"rsrvFee"`
		RsrvFeeExmpCode    interface{} `json:"rsrvFeeExmpCode"`
		SelectedVolume     interface{} `json:"selectedVolume"`
		PostDelvyYn        interface{} `json:"postDelvyYn"`
		CardGroupID        interface{} `json:"cardGroupId"`
		CardBpID           interface{} `json:"cardBpId"`
		SxTypeCode         interface{} `json:"sxTypeCode"`
		SxPrice            int         `json:"sxPrice"`
		SeatGradeName      string      `json:"seatGradeName"`
		AppSeatYn          interface{} `json:"appSeatYn"`
		PrintOrder         int         `json:"printOrder"`
		GradeColorVal      interface{} `json:"gradeColorVal"`
		UseYn              interface{} `json:"useYn"`
		SeatCount          int         `json:"seatCount"`
		ProdTicketTypeList []struct {
			ProdID           int         `json:"prodId"`
			ScheduleNo       string      `json:"scheduleNo"`
			SeatID           interface{} `json:"seatId"`
			FlplanID         int         `json:"flplanId"`
			SeatGradeNo      int         `json:"seatGradeNo"`
			SeatGradeName    string      `json:"seatGradeName"`
			BasePrice        int         `json:"basePrice"`
			GradeColorVal    interface{} `json:"gradeColorVal"`
			UseYn            interface{} `json:"useYn"`
			PriceNo          int         `json:"priceNo"`
			PriceName        string      `json:"priceName"`
			DcPriceTypeCode  string      `json:"dcPriceTypeCode"`
			DcTypeCode       string      `json:"dcTypeCode"`
			VolumeTypeCode   interface{} `json:"volumeTypeCode"`
			MaxSellVolume    int         `json:"maxSellVolume"`
			MinSellVolume    int         `json:"minSellVolume"`
			DcContent        string      `json:"dcContent"`
			BuyUnitVolume    int         `json:"buyUnitVolume"`
			PrevCondPriceNo  int         `json:"prevCondPriceNo"`
			PartCancelYn     string      `json:"partCancelYn"`
			DupDcYn          string      `json:"dupDcYn"`
			PostDelvyYn      string      `json:"postDelvyYn"`
			DcYn             string      `json:"dcYn"`
			PrintOrder       int         `json:"printOrder"`
			CardGroupID      interface{} `json:"cardGroupId"`
			CardBpID         interface{} `json:"cardBpId"`
			Mid              interface{} `json:"mid"`
			BasePriceDpYn    interface{} `json:"basePriceDpYn"`
			BpDcPriceDpYn    interface{} `json:"bpDcPriceDpYn"`
			LimitCnt         int         `json:"limitCnt"`
			SeatCount        int         `json:"seatCount"`
			TicketTypePrice  int         `json:"ticketTypePrice"`
			BasePriceYn      string      `json:"basePriceYn"`
			SelectedVolume   int         `json:"selectedVolume"`
			DpContent        interface{} `json:"dpContent"`
			KrPriceName      string      `json:"krPriceName"`
			DcValue          interface{} `json:"dcValue"`
			SejongPriceCode  interface{} `json:"sejongPriceCode"`
			DcPriceSoldOutYn string      `json:"dcPriceSoldOutYn"`
			DcPriceSellYCnt  int         `json:"dcPriceSellYCnt"`
			DispOrder        interface{} `json:"dispOrder"`
		} `json:"prodTicketTypeList"`
	} `json:"seatGradeList"`
	SeatGrades []string    `json:"seatGrades"`
	HTTPDomain interface{} `json:"httpDomain"`
}
type SaveHandler struct {
	FlplanTypeCode            string      `json:"flplanTypeCode"`
	Code                      string      `json:"code"`
	SeatInfoListWithPriceType string      `json:"seatInfoListWithPriceType"`
	CardCode                  string      `json:"cardCode"`
	Jtype                     string      `json:"jtype"`
	EType                     string      `json:"eType"`
	CustIP                    string      `json:"cust_ip"`
	ProdID                    string      `json:"prodId"`
	KakaoPayType              string      `json:"kakaoPayType"`
	UserName                  string      `json:"userName"`
	PayAmt                    string      `json:"payAmt"`
	PerfMainName              string      `json:"perfMainName"`
	PayNo                     string      `json:"payNo"`
	MidOptionKey              string      `json:"midOptionKey"`
	StaticDomain              interface{} `json:"staticDomain"`
	HTTPSDomain               interface{} `json:"httpsDomain"`
	Quota                     string      `json:"quota"`
	Tel                       string      `json:"tel"`
	RsrvSeq                   string      `json:"rsrvSeq"`
	PayMethodCode             string      `json:"payMethodCode"`
	HTTPDomain                interface{} `json:"httpDomain"`
}
