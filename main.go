package main

import (
	"net/http"
	"sea/config"
	"sea/infra"
	"sea/zlog"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	zlog.InitLogger("./log/Recommand.log", "debug")
	zlog.L().Info("service started")
	defer zlog.Sync()

	err := config.Load("./config.yaml")
	if err != nil {
		zlog.L().Error("config load failed",
			zap.Error(err))
		panic(err)
	}
	err = infra.MilvusInit()
	if err != nil {
		zlog.L().Error("milvus init failed",
			zap.Error(err))
		panic(err)
	}
	err = infra.Neo4jInit()
	if err != nil {
		zlog.L().Error("neo4j init failed",
			zap.Error(err))
		panic(err)
	}
	// 临时这么写，之后改
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	if err := router.Run(); err != nil {
		zlog.L().Error("http server run failed", zap.Error(err))
		panic(err)
	}
}
