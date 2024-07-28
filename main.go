package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"io"
	"math/rand/v2"
	"net/http"
	"os"
	"time"
)

var Version = "development"

const LogFormat = "%s | %s\n"

func main() {
	cacheEnabled := os.Getenv("ENABLE_CACHE") == "1"

	c := cache.New(60*time.Minute, 120*time.Minute)
	mux := http.NewServeMux()
	mux.HandleFunc("/write", func(writer http.ResponseWriter, request *http.Request) {
		data, _ := io.ReadAll(request.Body)
		if cacheEnabled {
			key := fmt.Sprintf("%v-%-4v", time.Now().UnixNano(), rand.IntN(9999))
			c.Set(key, data, cache.NoExpiration)
		}

		fmt.Printf("%s", data)
	})

	mux.HandleFunc("/readAll", func(writer http.ResponseWriter, request *http.Request) {
		if !cacheEnabled {
			writer.WriteHeader(http.StatusNotFound)
		}

		items := c.Items()
		for key, item := range items {
			_, _ = fmt.Fprintf(writer, LogFormat, key, item.Object)
		}
	})

	mux.HandleFunc("/health", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("OK"))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Logger Server (%s) is running on http://0.0.0.0:%s\n", Version, port)
	err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", port), mux)
	if err != nil {
		panic(err)
	}
}
