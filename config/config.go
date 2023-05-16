package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment               string // develop, staging, production
	CtxTimeout                int    // context timeout in seconds
	LogLevel                  string
	HTTPPort                  string
	SignInKey                 string
	AuthConfigPath            string
	CSVFilePath               string
	AccessTokenTimeOut        int
	PostgresHost              string
	PostgresPort              string
	PostgresDatabase          string
	PostgresUser              string
	PostgresPassword          string
	PostgresConnectionTimeOut int // seconds
	PostgresConnectionTry     int
}

// Load loads environment vars and inflates Config
func Load() Config {
	dotenvFilePath := cast.ToString(GetOrReturnDefault("DOT_ENV_PATH", "config/test.env"))
	err := godotenv.Load(dotenvFilePath)

	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	c := Config{}

	c.Environment = cast.ToString(GetOrReturnDefault("ENVIRONMENT", "develop"))
	c.LogLevel = cast.ToString(GetOrReturnDefault("LOG_LEVEL", "debug"))
	c.HTTPPort = cast.ToString(GetOrReturnDefault("HTTP_PORT", "8000"))
	c.SignInKey = cast.ToString(GetOrReturnDefault("SIGNINGKEY", "AzizbekSignIn"))
	c.CtxTimeout = cast.ToInt(GetOrReturnDefault("CTX_TIMEOUT", 7))
	c.AuthConfigPath = cast.ToString(GetOrReturnDefault("AUTH_PATH", "./config/auth.conf"))
	c.CSVFilePath = cast.ToString(GetOrReturnDefault("CSV_FILE_PATH", "./config/auth.csv"))
	c.AccessTokenTimeOut = cast.ToInt(GetOrReturnDefault("ACCESS_TOKEN_TIME_OUT", 5000))

	// Postgres
	c.PostgresHost = cast.ToString(GetOrReturnDefault("POSTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToString(GetOrReturnDefault("POSTGRES_PORT", 5432))
	c.PostgresDatabase = cast.ToString(GetOrReturnDefault("POSTGRES_DATABASE", "my_resume"))
	c.PostgresUser = cast.ToString(GetOrReturnDefault("POSTGRES_USER", "azizbek"))
	c.PostgresPassword = cast.ToString(GetOrReturnDefault("POSTGRES_PASSWORD", "Azizbek"))
	c.PostgresConnectionTimeOut = cast.ToInt(GetOrReturnDefault("POSTGRES_CONNECTION_TIMEOUT", 5))
	c.PostgresConnectionTry = cast.ToInt(GetOrReturnDefault("POSTGRES_CONNECTION_TRY", 10))

	return c
}

func GetOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
