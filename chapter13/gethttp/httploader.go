package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type channelInfo struct {
	url  string
	size int
}

func main() {
	myChannel := make(chan channelInfo)
	urls := []string{"https://example.com", "https://golang.org", "https://golang.org/doc"}
	for _, url := range urls {
		go responseSize(url, myChannel)
	}
	for i := 0; i < len(urls); i++ {
		page := <-myChannel
		fmt.Printf("%s: %d\n", page.url, page.size)
	}
}

func responseSize(url string, myChannel chan channelInfo) {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	result := channelInfo{url: url, size: len(body)}
	myChannel <- result
}
