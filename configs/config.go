package configs

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

type AppConfig struct {
	Port     int
	Driver   string
	Name     string
	Address  string
	DB_Port  int
	Username string
	Password string
	Secret string
	TEST_DB_Port int
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {
	var defaultConfig AppConfig
	defaultConfig.Port = getEnvInt("PORT", 9091)
	defaultConfig.Driver = getEnv("DRIVER", "mysql")
	defaultConfig.Name = getEnv("NAME", "magic_kodung")
	defaultConfig.Address = getEnv("ADDRESS", "localhost")
	defaultConfig.DB_Port = getEnvInt("DB_PORT",9876)
	defaultConfig.Username = getEnv("USERNAME", "root")
	defaultConfig.Password = getEnv("PASSWORD", "password")
	defaultConfig.Secret = getEnv("SECRET", "thisis32bitlongpassphraseimusing")
	defaultConfig.TEST_DB_Port = getEnvInt("TEST_DB_PORT",9987)

	fmt.Println(defaultConfig)

	return &defaultConfig
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		fmt.Println(value)
		return value
	}

	return fallback
}

func getEnvInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		fmt.Println(value)
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return intValue
		}
	}

	return fallback
}