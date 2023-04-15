package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/tsaqifammar/url-crawler/lib"
)

func main() {
	godotenv.Load()

	http.HandleFunc("/crawl", crawlHandler)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("A port must be set up")
	}

	log.Println("Starting web server at port", port)
	err := http.ListenAndServe("0.0.0.0:"+port, nil)
	if errors.Is(err, http.ErrServerClosed) {
		log.Println("Server closed")
	} else if err != nil {
		log.Println("Error starting the server", err)
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
