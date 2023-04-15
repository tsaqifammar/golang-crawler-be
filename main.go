package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/tsaqifammar/url-crawler/lib"
)

func main() {
	http.HandleFunc("/crawl", crawlHandler)

	fmt.Println("Starting web server at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server closed")
	} else if err != nil {
		fmt.Println("Error starting the server", err)
		os.Exit(1)
	}
}

func crawlHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "GET" {
		url := r.FormValue("url")
		depth := r.FormValue("depth")
		maxUrl := r.FormValue("maxUrl")

		if url == "" || depth == "" || maxUrl == "" {
			http.Error(w, "The parameters url, depth, and maxUrl are required", http.StatusBadRequest)
			return
		}

		if !lib.IsUrl(url) {
			http.Error(w, "The parameter url must be a valid url", http.StatusBadRequest)
			return
		}

		depthInt, err := strconv.Atoi(depth)
		if err != nil {
			http.Error(w, "The parameter depth must be an integer", http.StatusBadRequest)
			return
		}

		maxUrlInt, err := strconv.Atoi(maxUrl)
		if err != nil {
			http.Error(w, "The parameter maxUrl must be an integer", http.StatusBadRequest)
			return
		}

		c := lib.NewCrawler(url, depthInt, maxUrlInt)
		c.Crawl()

		res, err := json.Marshal(c.GetResults())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(res)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}
