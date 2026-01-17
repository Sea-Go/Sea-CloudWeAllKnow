package config

type Config struct {
	Milvus MilvusConfig `yaml:"milvus"`
}

type MilvusConfig struct {
	Address  string `yaml:"address"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}
