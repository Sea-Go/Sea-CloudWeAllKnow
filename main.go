package main

import (
	"sea/infra"
	"sea/zlog"

	"go.uber.org/zap"
)

func main() {
	zlog.InitLogger("./log/Recommand.log", "debug")
	zlog.L().Info("service started")

	err := infra.Milvus_Init()
	if err != nil {
		zlog.L().Error("infra init failed",
			zap.Error(err))
		panic(err)
	}

}
