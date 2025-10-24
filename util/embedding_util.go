package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const similarity_url = "http://localhost:8080/similarity"

const batch_similarity_url = "http://localhost:8080/batch_similarity"

// POST
const emgedding_url = "http://localhost:8080/embed"

type EmbeddingRequest struct {
	Text string `json:"text"`
}

type SimilarityRequest struct {
	Text1 string `json:"text1"`
	Text2 string `json:"text2"`
}

type BatchSimilarityRequest struct {
	TextA string   `json:"textA"`
	Texts []string `json:"texts"`
}

type EmbeddingResp struct {
	Text      string    `json:"text"`
	Embedding []float64 `json:"embedding"`
	Dimension int       `json:"dimension"`
	Count     int       `json:"count"`
}

type SimilarityResp struct {
	Text1      string  `json:"text1"`
	Text2      string  `json:"text2"`
	Similarity float64 `json:"similarity"`
}

type BatchSimilarityResp struct {
	TextA         string                    `json:"texta"`
	TotalCount    int                       `json:"total_count"`
	Similarities  []BatchSimilarityRespData `json:"similarities"`
	MaxSimilarity float64                   `json:"max_similarity"`
	MinSimilarity float64                   `json:"min_similarity"`
}

type BatchSimilarityRespData struct {
	Index      int     `json:"index"`
	Text       string  `json:"text"`
	Similarity float64 `json:"similarity"`
}

func GetEmbedding(text string) ([]float64, error) {
	log.Printf("GetEmbedding(%s)", text)
	//通过http调用 emgedding_url 用EmbeddingResp接收结果
	request := EmbeddingRequest{
		Text: text,
	}
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Post(emgedding_url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	log.Println("Embedding服务返回状态码:", resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		// 读取并返回服务的错误信息
		errorBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("服务返回非成功状态码 %d: %s", resp.StatusCode, string(errorBytes))
	}

	bodyBytes, _ := io.ReadAll(resp.Body)
	//log.Printf("body:%s", string(bodyBytes))

	// 5. 解析响应体
	var embeddingResp EmbeddingResp
	if err := json.Unmarshal(bodyBytes, &embeddingResp); err != nil {
		return nil, fmt.Errorf("解析 Embedding 服务响应失败: %w", err)
	}
	log.Printf("Embedding 服务响应 length:%d", len(embeddingResp.Embedding))
	return embeddingResp.Embedding, nil
}

func GetSimilarity(text1 string, text2 string) (float64, error) {
	log.Printf("GetSimilarity(%s,%s)", text1, text2)
	//通过http调用 emgedding_url 用EmbeddingResp接收结果
	request := SimilarityRequest{
		Text1: text1,
		Text2: text2,
	}
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return 0, err
	}

	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Post(similarity_url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	log.Println("similarity服务返回状态码:", resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		// 读取并返回服务的错误信息
		errorBytes, _ := io.ReadAll(resp.Body)
		return 0, fmt.Errorf("similarity服务返回非成功状态码 %d: %s", resp.StatusCode, string(errorBytes))
	}

	// 5. 解析响应体
	var sResp SimilarityResp

	if err := json.NewDecoder(resp.Body).Decode(&sResp); err != nil {
		return 0, fmt.Errorf("解析 similarity 服务响应失败: %w", err)
	}
	return sResp.Similarity, nil
}

func BatchGetSimilarity(text1 string, texts []string) ([]BatchSimilarityRespData, error) {
	//通过http调用 emgedding_url 用EmbeddingResp接收结果
	request := BatchSimilarityRequest{
		TextA: text1,
		Texts: texts,
	}
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Post(batch_similarity_url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	log.Println("similarity服务返回状态码:", resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		// 读取并返回服务的错误信息
		errorBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("similarity服务返回非成功状态码 %d: %s", resp.StatusCode, string(errorBytes))
	}
	bodyBytes, _ := io.ReadAll(resp.Body)
	log.Printf("body:%s", string(bodyBytes))
	// 5. 解析响应体
	var sResp BatchSimilarityResp

	if err := json.Unmarshal(bodyBytes, &sResp); err != nil {
		return nil, fmt.Errorf("解析 similarity 服务响应失败: %w", err)
	}
	return sResp.Similarities, nil
}
