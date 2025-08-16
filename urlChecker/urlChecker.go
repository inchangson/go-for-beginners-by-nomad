package urlchecker

import (
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

func Test() {
	results := make(map[string]string)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	ch := make(chan requestResult)

	for _, url := range urls {
		go hitURL(url, ch)
	}

	for i := 0; i < len(urls); i++ {
		res := <-ch
		results[res.url] = res.status
	}

	for url, status := range results {
		fmt.Printf("URL: %s, Status: %s\n", url, status)
	}

}

func hitURL(url string, ch chan<- requestResult) {
	// ch chan<- result로 정의하면 send-only 채널로 정의됨
	fmt.Println("Requesting:", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		ch <- requestResult{url: url, status: "failed"}
	} else {
		ch <- requestResult{url: url, status: "ok"}
	}
}
