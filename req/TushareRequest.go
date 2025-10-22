package req

type TushareRequest struct {
	APIName string                 `json: "apiName"`
	Token   string                 `json: "token"`
	Params  map[string]interface{} `json: params`
	Fields  string                 `json: fields,omitempty`
}
