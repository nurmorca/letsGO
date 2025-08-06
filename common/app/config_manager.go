package app

import "letsgo/common/postgresql"

type ConfigManager struct {
	PostgreSqlConfig postgresql.Config
}

func NewConfigManager() *ConfigManager {
	postgresConfig := getPostgreSqlConfig()
	return &ConfigManager{
		PostgreSqlConfig: postgresConfig,
	}
}

func getPostgreSqlConfig() postgresql.Config {
	return postgresql.Config{
		Host:                  "localhost",
		Port:                  "6432",
		Username:              "postgres",
		Password:              "postgres",
		DBname:                "productapp",
		MaxConnection:         "10",
		MaxConnectionIdleTime: "30s",
	}
}
