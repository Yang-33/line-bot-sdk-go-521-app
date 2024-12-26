package main

import (
	"fmt"
	"log"

	gohttplib "github.com/Yang-33/line-bot-sdk-go-521-lib"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: true, // disable adding timestamp for main_test.go
	})
	logrus.Info("Hello from logrus!")

	// special
	gohttplib.DumpAllDeps = true

	client := gohttplib.NewHttpClient()

	resp, err := client.Get("https://www.example.com")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println("Status:", resp.Status)
	// fmt.Println("Response Body:", string(body[:100]))
}
