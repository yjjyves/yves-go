package util

import (
	"context"
	"encoding/json"
	"log"
	"strings"
	"yves-go/entity"
)
import "github.com/elastic/go-elasticsearch/v8"

var es *elasticsearch.Client

func init() {
	log.Println("init elasticsearch.....")
	var err error
	es, err = elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://10.202.244.31:9200"},
		Username:  "elastic",
		Password:  "6Uu2EAcJ5DPG2t2T094zd4v4",
	})
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
}

func Search(language string) []*entity.NewsDocument {
	query := `{
        "query": {
            "bool": {
                "must": [
                    {"multi_match": {"query": "business", "fields": ["title","content"]}}
                ]
            }
        },
        "sort": [{"published_at": {"order": "desc"}}],
        "from": 0,
        "size": 10
    }`

	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("ai_news_en"),
		es.Search.WithBody(strings.NewReader(query)),
	)
	if err != nil {
		log.Printf("Error getting response: %s", err)
		return nil
	}
	defer func() { _ = res.Body.Close() }()

	var r struct {
		Hits struct {
			Hits []struct {
				Source entity.NewsDocument `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Printf("Error parsing response: %s", err)
		return nil
	}

	var newsList []*entity.NewsDocument
	for _, hit := range r.Hits.Hits {
		n := hit.Source
		newsList = append(newsList, &n)
	}

	return newsList
}
