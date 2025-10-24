package service

import (
	"context"
	"encoding/json"
	"log"
	"strings"
	"sync"
	"yves-go/entity"
	"yves-go/req"
	"yves-go/resp"
	"yves-go/util"

	"github.com/jinzhu/copier"
)

var EsNewsSearchServiceInstance = &EsNewsSearchService{}

type EsNewsSearchService struct {
}

func (esNewsSearchService *EsNewsSearchService) EsNewsSearch(req *req.NewsQueryReqVO) []*resp.NewsQueryRespVO {
	b, _ := json.Marshal(req)
	log.Printf("EsNewsSearchService.EsNewsSearch param: %s", string(b))
	res := Search(req.Language, "", "", "")
	return ConvertDocumentsToRespVO(res)
}

// ConvertDocumentsToRespVO 使用 copier 一行实现数组转换
func ConvertDocumentsToRespVO(docs []*entity.NewsDocument) []*resp.NewsQueryRespVO {
	if len(docs) == 0 {
		return nil
	}

	respList := make([]*resp.NewsQueryRespVO, len(docs))
	for i, doc := range docs {
		respList[i] = &resp.NewsQueryRespVO{}
		_ = copier.Copy(respList[i], doc) // 自动拷贝同名字段
	}

	return respList
}

func SearchV2(index string, query string) []*entity.NewsDocument {
	header := map[string]string{
		"Content-Type": "application/json",
	}

	res, err := util.Client.Search(
		util.Client.Search.WithContext(context.Background()),
		util.Client.Search.WithIndex(index),
		util.Client.Search.WithBody(strings.NewReader(query)),
		util.Client.Search.WithHeader(header),
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
				Score  float64             `json:"_score"`
				Id     string              `json:"_id"`
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
		n.Score = hit.Score
		n.Id = hit.Id
		newsList = append(newsList, &n)
	}

	return newsList
}

func Search(index string, keyword string, title string, content string) []*entity.NewsDocument {
	//query := `{
	//    "query": {
	//        "match_all": {}
	//    },
	//    "from": 0,
	//    "size": 5
	//}`

	query, _ := GenerateDynamicQuery(keyword, title, content)
	log.Println(query)

	return SearchV2(index, query)
}

// GenerateDynamicQuery 动态生成 Elasticsearch 查询体
func GenerateDynamicQuery(keyword string, title string, content string) (string, error) {
	// 存储所有需要满足的匹配条件
	mustClauses := make([]map[string]interface{}, 0)

	// 1. 检查 Title 关键词
	if keyword := strings.TrimSpace(keyword); keyword != "" {
		mustClauses = append(mustClauses, map[string]interface{}{
			"match": map[string]interface{}{
				"title": keyword,
			},
		})
	}

	// 2. 检查 Content 关键词
	if content := strings.TrimSpace(content); content != "" {
		mustClauses = append(mustClauses, map[string]interface{}{
			"match": map[string]interface{}{
				"content": content,
			},
		})
	}

	// 3. 检查 Keywords 关键词 (假设字段名为 tags)
	if title := strings.TrimSpace(title); title != "" {
		// 注意: 对于 tags/keywords 字段，如果希望精确匹配，可能使用 term 或 terms
		// 如果希望分词搜索，则使用 match
		mustClauses = append(mustClauses, map[string]interface{}{
			"match": map[string]interface{}{
				"title": title, // 假设字段名为 tags
			},
		})
	}

	queryBody := map[string]interface{}{
		"from": 0,
		"size": 5,
	}

	if len(mustClauses) > 0 {
		// 如果存在任何有效的关键词，则使用 bool/must
		queryBody["query"] = map[string]interface{}{
			"bool": map[string]interface{}{
				"must": mustClauses, // 逻辑 AND：必须同时满足所有非空的条件
			},
		}
	} else {
		//如果所有关键词都为空，则使用 match_all 查全部
		queryBody["query"] = map[string]interface{}{
			"match_all": map[string]interface{}{},
		}
	}

	// 将 Go 结构体序列化为 JSON 字符串
	jsonBody, err := json.MarshalIndent(queryBody, "", "    ")
	if err != nil {
		return "", err
	}

	return string(jsonBody), nil
}

func GenerateEmbeddingQuery(query string) (string, error) {
	// kNN 查询体
	queryVector, err := util.GetEmbedding(query)
	if err != nil {
		log.Println("get embedding error:", err)
		return "", err
	}
	knnQuery := map[string]interface{}{
		"field":          "content_vector", // 存储标题向量的字段名
		"query_vector":   queryVector,      // all-MiniLM-L6-v2 生成的查询向量
		"k":              10,               // 最终返回的结果数量 (Nearest Neighbors)
		"num_candidates": 100,              // 搜索时考虑的文档数量 (影响准确性和速度)
		// 如果需要，可以在这里添加 'filter' 来进行混合搜索 (例如：先按关键词过滤，再进行向量搜索)
		// "filter": []interface{}{
		//     map[string]interface{}{"term": map[string]interface{}{"category": "tech"}},
		// },
	}

	queryBody := map[string]interface{}{
		"from": 0,
		"size": 5,
		"knn":  knnQuery,
	}

	// 将 Go 结构体序列化为 JSON 字符串
	jsonBody, err := json.MarshalIndent(queryBody, "", "    ")
	if err != nil {
		return "", err
	}

	return string(jsonBody), nil
}

func MultiSearch(indexName string, query string, keyword string, title string, content string) []*entity.NewsDocument {
	if strings.TrimSpace(query) == "" {
		return nil
	}
	var wg = sync.WaitGroup{}
	wg.Add(2)

	resultChan := make(chan []*entity.NewsDocument, 2)

	//构造一路查询
	go func() {
		defer wg.Done()
		bm25Query, _ := GenerateDynamicQuery(keyword, title, content)
		bm25Result := SearchV2(indexName, bm25Query)
		log.Printf("bm25 finished: %d", len(bm25Result))
		resultChan <- bm25Result
	}()

	//构造二路查询
	go func() {
		defer wg.Done()
		knnQuery, _ := GenerateEmbeddingQuery(query)
		knnResult := SearchV2(indexName, knnQuery)
		log.Printf("knn finished:%d", len(knnResult))
		resultChan <- knnResult
	}()

	// 确保在所有 goroutine 完成后关闭 channel
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// --- 3. 进行RRF融合排序,收集结果 ---
	var resultLists [][]*entity.NewsDocument
	for res := range resultChan {
		if len(res) > 0 { // 过滤掉 nil 或空的结果集
			resultLists = append(resultLists, res)
		}
	}

	if len(resultLists) == 0 {
		log.Println("No search results were returned.")
		return nil
	}

	firstRankedDocuments := util.PerformRRFSorting(resultLists)
	log.Printf("Final ranked documents size: %d", len(firstRankedDocuments))

	finalRankedDocuments := util.PerformCrossEncoderSorting(firstRankedDocuments, query)
	// 返回 Top 10 或您需要的数量
	if len(finalRankedDocuments) > 10 {
		return finalRankedDocuments[:5]
	}
	return finalRankedDocuments
}
