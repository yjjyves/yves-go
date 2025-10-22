package resp

type TushareData struct {
	Fields []string        `json:"fields"`
	Items  [][]interface{} `json:"items"`
}
