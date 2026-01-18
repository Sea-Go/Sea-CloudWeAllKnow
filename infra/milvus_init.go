package infra

import (
	"context"
	"sea/config"
	"strings"

	"github.com/milvus-io/milvus/client/v2/milvusclient"
)

func MilvusInit() error {
	ctx := context.Background()
	cfg := config.Cfg
	client, err := milvusclient.New(ctx, &milvusclient.ClientConfig{
		Address:  cfg.Milvus.Address,
		Username: cfg.Milvus.Username,
		Password: cfg.Milvus.Password,
	})
	if err != nil {
		return err
	}

	db := strings.TrimSpace(cfg.Milvus.DBName)
	if db == "" {
		db = "default"
	}

	if err := client.CreateDatabase(ctx, milvusclient.NewCreateDatabaseOption(db)); err != nil {
		// 已存在就忽略
		msg := strings.ToLower(err.Error())
		if !strings.Contains(msg, "already exists") && !strings.Contains(msg, "exist") {
			return err
		}
	}

	if err := client.UseDatabase(ctx, milvusclient.NewUseDatabaseOption(db)); err != nil {
		return err
	}
	
	return nil
}
