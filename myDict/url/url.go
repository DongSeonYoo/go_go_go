package url

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestfailed = errors.New("request failed")

func Url() {
	var results = map[string]string{}
	url := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/qwe",
		"https://academy.nomardcoders.co/",
	}

	for _, url := range url {
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAIL"
		}
		results[url] = result
	}
	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string) error {
	fmt.Println("Checking, ", url)
	res, err := http.Get(url)
	if err != nil || res.StatusCode >= 400 {
		return errRequestfailed
	}

	return nil
}
