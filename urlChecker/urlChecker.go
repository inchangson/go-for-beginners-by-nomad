package urlchecker

import (
	"fmt"
	"net/http"
)

func Test() {
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

	// map 초기화 방법 1
	// results := map[string]string{}
	// map 초기화 방법 2
	results := make(map[string]string)

	for _, url := range urls {
		err := hitURL(url)
		if err != nil {
			results[url] = "FAILED"
		} else {
			results[url] = "OK"
		}
	}

	for url, result := range results {
		fmt.Println(url, ":", result)
	}
}

var errRequestFailed error = fmt.Errorf("request failed")

func hitURL(url string) error {
	fmt.Println("Requesting:", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println("Error Status Code:", resp.StatusCode)
		return errRequestFailed
	}
	return nil
}
