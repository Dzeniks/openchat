// main.go
package main

import (
	"log"
	"os"
	env "server-go/lib/dotEnv"
	"server-go/lib/secretKey"

	"server-go/routes"

	"github.com/gin-gonic/gin"
)

func createLogFile(filename string) (*os.File, error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func main() {
	r := gin.Default()
	logFile, err := createLogFile("log.txt")
	if err != nil {
		log.Fatal("Unable to create log file:", err)
	}
	defer func(logFile *os.File) {
		err := logFile.Close()
		if err != nil {
			log.Fatal("Unable to close log file:", err)
		}
	}(logFile)
	log.Printf("Log file created: %s", logFile.Name())

	if env.DotEnv.Production == true {
		log.SetOutput(logFile)
	}

	//goland:noinspection Annotator
	err = env.LoadDotEnv()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DotEnv loaded")
	secretKey.SetSecretKey(env.DotEnv.SecretKey)

	log.Println(env.DotEnv.SecretKey)
	log.Println(secretKey.GetSecretKey())

	routes.InitApiRouter(r)

	err = r.Run()
	if err != nil {
		return
	}
}
