package app

import (
	"flag"
	"gotoko-postgres/app/controllers"
	"log"
	"os"

	"github.com/joho/godotenv"
)


func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func Run()  {
	var server = controllers.Server{}
	var appConfig = controllers.AppConfig{}
	var dbConfig = controllers.DBconfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error on loading .env file")
	}

	appConfig.AppName = getEnv("APP_NAME", "GoTokoWEB")
	appConfig.AppEnv = getEnv("APP_ENV", "Development")
	appConfig.AppPort = getEnv("APP_PORT", "9000")

	dbConfig.DBHost = getEnv("DB_HOST", "localhost")
	dbConfig.DBUser = getEnv("DB_USER", "yosaae24")
	dbConfig.DBPassword = getEnv("DB_PASSWORD", "minusyosa24")
	dbConfig.DBPort = getEnv("DB_PORT", "5432")
	dbConfig.DBName = getEnv("DB_NAME", "gotoko-postgres")

	flag.Parse()
	arg := flag.Arg(0)
	if arg != "" {
		server.InitCommands(appConfig, dbConfig)
	} else {
		server.Initialize(appConfig, dbConfig)
		server.Run(":" + appConfig.AppPort)
	}

}