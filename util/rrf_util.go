package util

import (
	"sort"
	"yves-go/entity"
)

// RRFDocument 结构体用于 RRF 融合计算
type RRFDocument struct {
	Document *entity.NewsDocument
	RRFScore float64
}

// RRF 算法的平滑常数 K
const K = 60.0

// PerformRRFSorting 对多路搜索结果进行 RRF 融合排序
// inputs: 多个 []*entity.NewsDocument 列表 (如 BM25 和 KNN 结果)
// returns: 融合排序后的 []*entity.NewsDocument
func PerformRRFSorting(resultLists [][]*entity.NewsDocument) []*entity.NewsDocument {
	// 1. 构建文档ID到RRFDocument的映射
	// 使用 map[string]*RRFDocument 确保每个文档只被计算一次
	documentScores := make(map[string]*RRFDocument)

	// 2. 遍历所有结果列表（多路查询）
	for _, resultList := range resultLists {
		// 3. 遍历列表中的每个文档，计算其 RRF 贡献分
		for rank, doc := range resultList {
			// Rank 是从 0 开始的索引，RRF公式中的 Rank 是从 1 开始
			rrfRank := float64(rank + 1)

			// 计算当前文档在当前列表中的 RRF 贡献分数
			contribution := 1.0 / (K + rrfRank)

			// 如果文档已经存在于 Map 中，累加分数
			if rrfDoc, ok := documentScores[doc.Id]; ok {
				rrfDoc.RRFScore += contribution
			} else {
				// 如果是新文档，初始化并设置初始分数
				documentScores[doc.Id] = &RRFDocument{
					Document: doc,
					RRFScore: contribution,
				}
			}
		}
	}

	// 4. 将 Map 中的 RRFDocument 转换为切片以便排序
	finalRankedList := make([]*RRFDocument, 0, len(documentScores))
	for _, rrfDoc := range documentScores {
		finalRankedList = append(finalRankedList, rrfDoc)
	}

	// 5. 按 RRFScore 降序排序
	sort.Slice(finalRankedList, func(i, j int) bool {
		return finalRankedList[i].RRFScore > finalRankedList[j].RRFScore
	})

	// 6. 提取并返回排序后的原始文档列表
	finalDocuments := make([]*entity.NewsDocument, 0, len(finalRankedList))
	for _, rrfDoc := range finalRankedList {
		finalDocuments = append(finalDocuments, rrfDoc.Document)
	}

	return finalDocuments
}

// 进行cross-encoder模型排序
func PerformCrossEncoderSorting(resultLists []*entity.NewsDocument, query string) []*entity.NewsDocument {
	// 1. 构建文档ID到RRFDocument的映射
	// 使用 map[string]*RRFDocument 确保每个文档只被计算一次
	documentScores := make(map[string]*RRFDocument)

	texts := make([]string, 0)

	//获取所有的text
	for _, doc := range resultLists {
		// 如果文档已经存在于 Map 中，累加分数
		if _, ok := documentScores[doc.Content]; !ok {
			texts = append(texts, doc.Content)
			documentScores[doc.Content] = &RRFDocument{
				Document: doc,
				RRFScore: 0,
			}
		}
	}

	//获取相似度
	similarity, err := BatchGetSimilarity(query, texts)
	if err != nil {
		return nil
	}

	// 4. 将 Map 中的 RRFDocument 设置相似度以便排序
	finalRankedList := make([]*RRFDocument, 0, len(documentScores))

	for _, slt := range similarity {
		if _, ok := documentScores[slt.Text]; ok {
			//根据slt.Text获取RRFDocument
			rrfDoc := documentScores[slt.Text]
			rrfDoc.RRFScore = slt.Similarity
			rrfDoc.Document.Score = slt.Similarity
			finalRankedList = append(finalRankedList, rrfDoc)
		}
	}

	// 5. 按 RRFScore 降序排序
	sort.Slice(finalRankedList, func(i, j int) bool {
		return finalRankedList[i].RRFScore > finalRankedList[j].RRFScore
	})

	// 6. 提取并返回排序后的原始文档列表
	finalDocuments := make([]*entity.NewsDocument, 0, len(finalRankedList))
	for _, rrfDoc := range finalRankedList {
		finalDocuments = append(finalDocuments, rrfDoc.Document)
	}

	return finalDocuments
}
