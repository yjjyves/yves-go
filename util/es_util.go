package util

import (
	"log"
)
import "github.com/elastic/go-elasticsearch/v8"

var Client *elasticsearch.Client

func init() {
	log.Println("init elasticsearch.....")
	var err error
	Client, err = elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://10.202.244.31:9200"},
		Username:  "elastic",
		Password:  "6Uu2EAcJ5DPG2t2T094zd4v4",
	})
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
}
