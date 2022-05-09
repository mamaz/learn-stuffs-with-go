package external

import (
	"net/http"
)

// WhichFaster will return which are faster between 2 URLs
func WhichFaster(firstURL, secondURL string) string {
	// select will check multiple channels wich one's returning value first
	select {
	case res := <-ping(firstURL):
		return res
	case res := <-ping(secondURL):
		return res
	}
}

func ping(url string) chan string {
	channel := make(chan string)

	go func(ch chan string) {
		http.Get(url)

		ch <- url

		close(ch)
	}(channel)

	return channel
}
