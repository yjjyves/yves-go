package test

import (
	"encoding/json"
	"log"
	"testing"
	"time"
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

func TestEsQuery02(t *testing.T) {
	log.Println("====")

	start := time.Now()
	result := service.MultiSearch("ai_news_en_vec",
		"What awards did Mitsubishi Elevator Thailand win at the Asia Pacific Enterprise Awards 2025?", "",
		"", "Elevator")

	log.Println("耗时: ", time.Since(start))

	if data, err := json.MarshalIndent(result, "", "  "); err == nil {
		log.Println("result: ", string(data))
	} else {
		log.Println("marshal error:", err)
	}
}
