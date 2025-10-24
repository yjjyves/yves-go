package test

import (
	"encoding/json"
	"log"
	"testing"
	"yves-go/service"
)

func TestEsQuery01(t *testing.T) {
	log.Println("====")

	result := service.Search("ai_news_en_vec", "", "Stanbroke", "The expanding fine dining sector and evolving consumer preferences")

	if data, err := json.MarshalIndent(result, "", "  "); err == nil {
		log.Println("result: ", string(data))
	} else {
		log.Println("marshal error:", err)
	}
}
