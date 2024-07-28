package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"io"
	"math/rand/v2"
	"net/http"
	"os"
	"slices"
	"sort"
	"time"
)

var Version = "development"

const LogFormat = "%s | %s\n"

func main() {
	preserveLogEnabled := os.Getenv("PRESERVE_LOG_ENABLED") == "1"

	c := cache.New(60*time.Minute, 120*time.Minute)
	mux := http.NewServeMux()
	mux.HandleFunc("/write", func(writer http.ResponseWriter, request *http.Request) {
		data, _ := io.ReadAll(request.Body)
		if preserveLogEnabled {
			key := fmt.Sprintf("%v-%-4v", time.Now().UnixNano(), rand.IntN(9999))
			c.Set(key, data, cache.NoExpiration)
		}

		fmt.Printf("%s\n", data)
	})

	mux.HandleFunc("/readAll", func(writer http.ResponseWriter, request *http.Request) {
		if !preserveLogEnabled {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		items := c.Items()
		keys := make([]string, 0)
		for k := range items {
			keys = append(keys, k)
		}
		slices.Sort(keys)
		sort.Strings(keys)

		for _, k := range keys {
			_, _ = fmt.Fprintf(writer, LogFormat, k, items[k].Object)
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
