package service

import (
	"encoding/json"
	"log"
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
	res := util.Search(req.Language)
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
