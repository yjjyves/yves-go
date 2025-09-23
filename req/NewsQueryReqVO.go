package req

type NewsQueryReqVO struct {
	Language string `json:"language"`
	Query    string `json:"query"`
}
