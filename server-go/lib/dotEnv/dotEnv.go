package dotEnv

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	HostPort           string
	MongoConnectionURI string
	MongoUsername      string
	MongoPassword      string
	MongoDatabase      string
	SecretKey          string
	Production         bool
	AI_URL             string
}

var DotEnv Env

var requiredEnv = []string{"HOST_PORT", "MONGO_CONNECTION_STRING", "MONGO_USERNAME", "MONGO_PASSWORD", "MONGO_DBNAME",
	"SECRET_KEY", "PRODUCTION", "AI_URL"}

func checkDotEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	for _, env := range requiredEnv {
		envVar := os.Getenv(env)
		if envVar == "" {
			return fmt.Errorf("environment variable %s not set", env)
		}
	}
	return nil
}

//goland:noinspection GoUnusedExportedFunction
func LoadDotEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	err = checkDotEnv()
	if err != nil {
		return err
	}
	DotEnv = Env{
		HostPort:           os.Getenv("HOST_PORT"),
		MongoConnectionURI: os.Getenv("MONGO_CONNECTION_STRING"),
		MongoUsername:      os.Getenv("MONGO_USERNAME"),
		MongoPassword:      os.Getenv("MONGO_PASSWORD"),
		MongoDatabase:      os.Getenv("MONGO_DBNAME"),
		SecretKey:          os.Getenv("SECRET_KEY"),
		Production:         os.Getenv("PRODUCTION") == "true",
		AI_URL:             os.Getenv("AI_URL"),
	}
	return nil
}
