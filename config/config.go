package config

import (
	"os"
	"sea/zlog"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

var Cfg Config

type Config struct {
	Milvus MilvusConfig `mapstructure:"milvus" yaml:"milvus"`
	Ali    AliConfig    `mapstructure:"ali" yaml:"ali"`
	Kafka  KafkaConfig  `mapstructure:"Kafka" yaml:"Kafka"` // note: key is "Kafka" in your YAML
	Neo4j  Neo4jConfig  `mapstructure:"neo4j" yaml:"neo4j"`
}

type MilvusConfig struct {
	Address  string `mapstructure:"address" yaml:"address"`
	Username string `mapstructure:"username" yaml:"username"`
	Password string `mapstructure:"password" yaml:"password"`
	DBName   string `mapstructure:"dbname" yaml:"dbname"`
}

type AliConfig struct {
	APIKey  string `mapstructure:"apikey" yaml:"apikey"`
	BaseURL string `mapstructure:"baseurl" yaml:"baseurl"`
}

type KafkaConfig struct {
	Address string `mapstructure:"address" yaml:"address"`
}

type Neo4jConfig struct {
	Address  string `mapstructure:"address" yaml:"address"`
	Username string `mapstructure:"username" yaml:"username"`
	Password string `mapstructure:"password" yaml:"password"`
}

func Load(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		zlog.L().Error("read config file error", zap.Error(err))
		return err
	}

	if err := yaml.Unmarshal(data, &Cfg); err != nil {
		zlog.L().Error("unmarshal config file error", zap.Error(err))
		return err
	}

	return nil
}
