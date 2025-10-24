package test

import (
	"log"
	"testing"
	"time"
	"yves-go/util"
)

func TestQueryEmbedding(t *testing.T) {
	currentTime := time.Now()
	embedding, err := util.GetEmbedding("The expanding fine dining sector and evolving consumer preferences")

	//打印耗时
	log.Println("耗时: ", time.Since(currentTime))

	if err != nil {
		log.Println("get embedding error:", err)
		return
	}

	log.Println("embedding: ", embedding)
}

func TestQuerySimilarity(t *testing.T) {
	currentTime := time.Now()
	similarity, err := util.GetSimilarity("The expanding fine dining sector and evolving consumer preferences", "Bangkok continues to impress with its ever-evolving food and beverage scene. New restaurants and drink spots are opening at a rapid pace, offering exciting culinary experiences. From Isan-inspired fusion restaurants like Somsak and Arloy to modern Asian fine dining at Samas and Mexican-Thai fusion at Macaria, Bangkok’s dining landscape is vibrant and diverse. Whether it’s the revival of traditional recipes or innovative twists on global cuisine, the city’s latest F&B offerings are a testament to its status as a global food capital.")
	//打印耗时
	log.Println("耗时: ", time.Since(currentTime))
	if err != nil {
		log.Println("get similarity error:", err)
		return
	}
	log.Println("similarity: ", similarity)

}

func TestBatchQuerySimilarity(t *testing.T) {
	currentTime := time.Now()

	testA := "The expanding fine dining sector and evolving consumer preferences"
	tests := []string{
		"Bangkok continues to impress with its ever-evolving food and beverage scene. New restaurants and drink spots are opening at a rapid pace, offering exciting culinary experiences. From Isan-inspired fusion restaurants like Somsak and Arloy to modern Asian fine dining at Samas and Mexican-Thai fusion at Macaria, Bangkok’s dining landscape is vibrant and diverse. Whether it’s the revival of traditional recipes or innovative twists on global cuisine, the city’s latest F&B offerings are a testament to its status as a global food capital",
		"The Thai shabu-shabu restaurant market is becoming increasingly competitive with the launch of Teenoi Gold, a new brand by BNN Restaurant Group targeting the premium mass segment. Industry experts warn that larger players may dominate the market, squeezing out smaller businesses. With rising competition and economic challenges, the new brand aims to differentiate itself through premium offerings and Cantonese-style cuisine, priced at 599 baht per person. The company plans to expand to 11 outlets by 2026",
		"Australia’s Stanbroke and Thailand’s Udom Supply have launched True North, a premium beef brand, " +
			"into Thailand’s market. The product was introduced at a launch event at the Bangkok Marriott Hotel Sukhumvit, where it was showcased to over 200 industry professionals. Udom Supply will be the exclusive distributor of the brand in Thailand, aiming to elevate the country’s premium dining standards. The partnership reflects growing demand for high-quality imported beef in Thailand, driven by the expanding fine dining sector and evolving consumer preferences.",
	}

	similarity, err := util.BatchGetSimilarity(testA, tests)
	//打印耗时
	log.Println("耗时: ", time.Since(currentTime))
	if err != nil {
		log.Println("get similarity error:", err)
		return
	}
	log.Println("similarity: ", similarity)

}
