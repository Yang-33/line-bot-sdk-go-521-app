package main

import (
	"fmt"
	"io"
	"log"

	gohttplib "github.com/Yang-33/line-bot-sdk-go-521-lib"
)

func main() {
	// special
	gohttplib.DumpAllDeps = true

	client := gohttplib.NewHttpClient()

	resp, err := client.Get("https://www.example.com")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Status:", resp.Status)
	fmt.Println("Response Body:", string(body[:100]))
}
