package external

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type URLResult struct {
	URL     string
	IsValid bool
	Reason  string
}

// CheckValidURLs will check which URL is valid between slice or urls
// valid URL is the one which return "valid": "true" value and return 200
func CheckValidURLs(urls []string, result chan URLResult) {
	for _, url := range urls {
		go fetchValidURLs(url, result)
	}
}

func fetchValidURLs(url string, result chan URLResult) {
	resp, err := http.Get(url)

	if err != nil {
		result <- URLResult{
			URL:     url,
			IsValid: false,
			Reason:  err.Error(),
		}
		return
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		result <- URLResult{
			URL:     url,
			IsValid: false,
			Reason:  err.Error(),
		}
		return
	}

	var data map[string]string
	err = json.Unmarshal(b, &data)
	if err != nil {
		result <- URLResult{
			URL:     url,
			IsValid: false,
			Reason:  err.Error(),
		}
		return
	}

	result <- URLResult{
		URL:     url,
		IsValid: data["valid"] == "true",
		Reason:  "",
	}
}
