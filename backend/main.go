package main

import (
	_ "backend/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type PingResponse struct {
    Message string `json:"message"`
}

// @title Gin APIドキュメント
// @version 1.0
// @description Ginフレームワークで作成したREST APIのドキュメント
// @host localhost:8080
// @BasePath /

func main() {
    r := gin.Default()

    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    r.GET("/ping", PingHandler)

    r.Run("0.0.0.0:8080")
}

// @Summary 動作確認
// @Description サーバーの稼働確認用API
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} PingResponse "正常応答"
// @Router /ping [get]
func PingHandler(c *gin.Context) {
    response := PingResponse{Message: "pong"}
    c.JSON(200, response)
}
