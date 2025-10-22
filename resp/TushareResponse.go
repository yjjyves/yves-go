package resp

import "yves-go/resp"

type TushareResponse struct {
	Code int              `json:"code"`
	Msg  string           `json:"msg"`
	Data resp.TushareData `json:data`
}
