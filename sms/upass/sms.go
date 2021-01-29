package upass

import (
	"hello/library/uhttp"
	"strings"
)

type uPass struct {
	Phone      string   `json:"phone"`
	Templateid string   `json:"templateid"`
	Params     []string `json:"params"`
}

const (
	url    = "https://open.ucpaas.com/ol/sms/sendsms"
	sid    = "d093d1804b46eeab43fe3be79923649f"
	app_id = "3bdc9fbb73924971ae963262e887157d"
	token  = "b9cc26265fbf1a83cef1d0c2b60cdc7f"
)

func NewUPass() *uPass {
	return &uPass{}
}

func (u *uPass) SendPass() map[string]interface{} {
	params := map[string]interface{}{
		"sid":    sid,
		"app_id": app_id,
		"token":  token,
	}
	params["templateid"] = u.Templateid
	params["param"] = strings.Join(u.Params, ",")
	params["mobile"] = u.Phone
	return uhttp.Post(url, params)
}
