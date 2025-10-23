package resp

type TushareResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data TushareData `json:data`
}
