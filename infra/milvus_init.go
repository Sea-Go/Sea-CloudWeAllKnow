package infra

import (
	"context"
	"fmt"
	"os"
	"sea/config"
	"strings"

	"github.com/milvus-io/milvus/client/v2/milvusclient"
	"gopkg.in/yaml.v3"
)

var (
	Milvus *milvusclient.Client
	Cfg    config.Config
)

func Milvus_Init() error {
	data, err := os.ReadFile("./config.yaml")
	if err != nil {
		return fmt.Errorf("read config.yaml: %w", err)
	}

	if err := yaml.Unmarshal(data, &Cfg); err != nil {
		return fmt.Errorf("parse config.yaml: %w", err)
	}

	ctx := context.Background()

	client, err := milvusclient.New(ctx, &milvusclient.ClientConfig{
		Address:  Cfg.Milvus.Address,
		Username: Cfg.Milvus.Username,
		Password: Cfg.Milvus.Password,
	})
	if err != nil {
		return fmt.Errorf("infra milvus client: %w", err)
	}

	db := strings.TrimSpace(Cfg.Milvus.DBName)
	if db == "" {
		db = "default"
	}

	if err := client.CreateDatabase(ctx, milvusclient.NewCreateDatabaseOption(db)); err != nil {
		// 已存在就忽略
		msg := strings.ToLower(err.Error())
		if !strings.Contains(msg, "already exists") && !strings.Contains(msg, "exist") {
			return fmt.Errorf("create database %q: %w", db, err)
		}
	}

	if err := client.UseDatabase(ctx, milvusclient.NewUseDatabaseOption(db)); err != nil {
		return fmt.Errorf("use database %q: %w", db, err)
	}

	Milvus = client
	return nil
}
