package tushare

import (
	"log"
	"yves-go/entity"

	"github.com/gin-gonic/gin"
)

const tushareAPIURL = "http://api.tushare.pro"
const tushareToken = "0c53281b41d5cce35a40d8f1185b1db8a13c724a1cfb3d6a31e4ba05" // 替换为你的Tushare Token

func RegisterTushareApi(r *gin.Engine) {
	log.Println("RegisterTushareApi /news/query")
	r.GET("/tushare", nil)
}

func GetDataV1() *entity.TushareResponse {
	return nil
}
