package model

type GetMember struct {
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
	Data    struct {
		MemberKey int    `json:"memberKey"`
		Email     string `json:"email"`
		Name      string `json:"name"`
	} `json:"data"`
}
