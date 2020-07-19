package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	// General
	AppName string `mapstructure:"appname"`
	Env     string `mapstructure:"env"`

	// Server
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`

	HttpClientTimeout int `mapstructure:"httpclient_timeout"`

	// Token
	Secret string `mapstructure:"secret"`

	// DB -> MySQL
	MysqlHost     string `mapstructure:"mysql_host"`
	MysqlPassword string `mapstructure:"mysql_password"`
	MysqlDB       string `mapstructure:"mysql_db"`
	MysqlUser     string `mapstructure:"mysql_user"`
	MysqlPort     string `mapstructure:"mysql_port"`

	// Caching -> Redis
	RedisHost string `mapstructure:"redis_host"`
	RedisPort string `mapstructure:"redis_port"`

	// MQ -> Google PubSub
	GoogleProjectID          string `mapstructure:"google_project_id"`
	GoogleClusterName        string `mapstructure:"google_cluster_name"`
	GoogleCredentialFilepath string `mapstructure:"google_credential_filepath"`
	GoogleMaxRequeueNotifier int    `mapstructure:"google_max_requeue_notifier"`
	GoogleMaxRequeueConsumer int    `mapstructure:"google_max_requeue_consumer"`
	GoogleConcurrent         int    `mapstructure:"google_concurrent"`
	GoogleMaxInFlight        int    `mapstructure:"google_max_in_flight"`

	// JWT
	JwtAccessExpires  int `mapstructure:"jwt_at_expire"`
	JwtRefreshExpires int `mapstructure:"jwt_rt_expire"`
}

func parseConfigFilePath() string {
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return filepath.Join(workPath, "config")
}

var AppConfig *Config

func init() {
	AppConfig = InitAppConfig()
}

func InitAppConfig() *Config {
	configPath := parseConfigFilePath()
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	config := new(Config)
	if err := viper.Unmarshal(config); err != nil {
		panic(fmt.Errorf("failed to parse config file: %w", err))
	}

	return config
}
