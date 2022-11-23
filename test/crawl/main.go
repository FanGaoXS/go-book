package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

var (
	urls = []string{"https://www.baidu.com/s"}
)

func main() {
	workList := make(chan []string)
	go func() {
		workList <- urls
	}()

	seen := make(map[string]struct{})
	for {
		list := <-workList
		for _, link := range list {
			if _, ok := seen[link]; ok {
				continue
			}
			seen[link] = struct{}{}
			go func(link string) {
				workList <- crawl(link)
			}(link)
		}
	}
}

func crawl(url string) []string {
	log.Printf("url is %s", url)
	links, err := fetch(url)
	if err != nil {
		log.Fatalln(err)
	}

	return links
}

func fetch(url string) ([]string, error) {
	time.Sleep(time.Second * 2)
	res := make([]string, 0, 0)
	rand.Seed(time.Now().Unix())
	for i := 0; i < rand.Intn(10); i++ {
		res = append(res, fmt.Sprintf("%s?query=%d", url, i+1))
	}
	log.Printf("get %d links from %s", len(res), url)
	return res, nil
}
