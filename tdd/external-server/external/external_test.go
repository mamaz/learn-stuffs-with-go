// This test will show external server can be created easily for
// testing purposes with httptest package
package external

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestExternalServer(t *testing.T) {
	t.Run("it should be able to check which one is faster", func(t *testing.T) {
		// arrange
		fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(1 * time.Millisecond)
			w.WriteHeader(http.StatusOK)
		}))

		slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(50 * time.Millisecond)
			w.WriteHeader(http.StatusOK)
		}))

		defer fastServer.Close()
		defer slowServer.Close()

		// act
		fasterURL := WhichFaster(fastServer.URL, slowServer.URL)

		// assert
		assert.Equal(t, fastServer.URL, fasterURL)
	})
}

func TestValidURL(t *testing.T) {
	t.Run("it should be able to check valid urls", func(t *testing.T) {
		serverOne := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(100 * time.Millisecond)
			w.Write([]byte(`{
				"valid": "true"
			}`))

		}))

		serverTwo := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(100 * time.Millisecond)
			w.Write([]byte(`{}`))
		}))

		serverThree := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(100 * time.Millisecond)
			w.Write([]byte(`{
				"valid": "false"
			}`))
		}))

		serverFour := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(100 * time.Millisecond)
			w.WriteHeader(http.StatusInternalServerError)
		}))

		defer serverOne.Close()
		defer serverTwo.Close()
		defer serverThree.Close()
		defer serverFour.Close()

		channel := make(chan URLResult)

		urls := []string{
			serverOne.URL,
			serverTwo.URL,
			serverThree.URL,
			serverFour.URL,
		}

		CheckValidURLs(urls, channel)

		// filter valid value
		result := []URLResult{}
		for i := 0; i < len(urls); i++ {
			urlresult := <-channel
			if urlresult.IsValid == true {
				result = append(result, urlresult)
			}
		}

		// it should have one valid url
		// the other 3 should be failed
		assert.Equal(t, 1, len(result))
	})
}
