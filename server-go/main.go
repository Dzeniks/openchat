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

	if gin.Mode() == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
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
	port := env.DotEnv.HostPort
	if port == "" {
		port = "8080"
	}

	r.ForwardedByClientIP = true
	err = r.SetTrustedProxies([]string{"127.0.0.1", "192.168.1.2", "10.0.0.0/8"})
	if err != nil {
		return
	}

	err = r.Run(":" + port)
	if err != nil {
		return
	}

}
