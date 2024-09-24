package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"time"
)

type URL struct {
	Id           string    `json:"Id"`
	OriginalUrl  string    `json:"OriginalUrl"`
	ShortUrl     string    `json:"ShortUrl"`
	CreationDate time.Time `json:"CreationDate"`
}

// in memory db

var urlDb = make(map[string]URL)

func genrateShortUrl(OriginalUrl string) string {
	// function that will convert into hash
	hasher := md5.New()
	hasher.Write([]byte(OriginalUrl))
	fmt.Println("hasher", hasher)
	data := hasher.Sum(nil)
	fmt.Println("hasher data:", data)
	hash := hex.EncodeToString(data)
	fmt.Println("encoded String:", hash)
	fmt.Println("final Sting:", hash[:8])
	return hash[:8]

}

// storing in db

func createURL(OriginalUrl string) string {
	shortUrl := genrateShortUrl(OriginalUrl)
	id := shortUrl
	urlDb[id] = URL{
		Id:           id,
		OriginalUrl:  OriginalUrl,
		ShortUrl:     shortUrl,
		CreationDate: time.Now(),
	}
	return shortUrl
}

func getURL(id string) (URL, error) {
	url, ok := urlDb[id]
	if !ok {
		return URL{}, errors.New("URL Not found")
	}
	return url, nil
}

func RootPageURL(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Worldss")
}

func ShortURLHandlers(w http.ResponseWriter, r *http.Request) {
	var data struct {
		URL string `json:"url"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadGateway)
		return
	}
	shortURL_ := createURL(data.URL)
	// fmt.Fprint(w, shortURL_)
	response := struct {
		ShortURL string `json:"short_url"`
	}{ShortURL: shortURL_}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func redirectURLHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/redirect/"):]
	url, err := getURL(id)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusNotFound)
	}
	http.Redirect(w, r, url.OriginalUrl, http.StatusFound)
}

func main() {
	// fmt.Println("we are creating url shortener ")
	// OriginalUrl := "https://github.com/palindrome-kasak"
	// genrateShortUrl(OriginalUrl)
	// Register the handler function to handle all requests to the root url("/")
	http.HandleFunc("/", RootPageURL)
	http.HandleFunc("/shorten", ShortURLHandlers)
	http.HandleFunc("/redirect/", redirectURLHandler)

	// setup server , local host
	// http server start
	fmt.Println("starting the http server on port 3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error on starting server:,", err)
	}
}
