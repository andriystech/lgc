package config

import (
	"os"
	"time"
)

type ServerConfig struct {
	Port                         string
	MongoDbUrl                   string
	DbName                       string
	DbConnectionTimeoutInSeconds int
	TokenTTLInSeconds            int
	WsReadBuffer                 int
	WsWriteBuffer                int
}

const defaultPort = ":8090"
const defaultMongoDbURL = "mongodb://root:123456@localhost:27017"
const defaultDbName = "lgc"

func GetServerConfig() *ServerConfig {

	return &ServerConfig{
		Port:                         env("SERVER_PORT", defaultPort),
		MongoDbUrl:                   env("MONGODB_URL", defaultMongoDbURL),
		DbName:                       env("MONGO_DB_NAME", defaultDbName),
		DbConnectionTimeoutInSeconds: int(time.Second * 20),
		TokenTTLInSeconds:            int(time.Second * 60),
		WsReadBuffer:                 1000,
		WsWriteBuffer:                1000,
	}
}

func env(variable string, defaultValue string) string {
	value, ok := os.LookupEnv(variable)
	if !ok {
		value = defaultValue
	}
	return value
}
