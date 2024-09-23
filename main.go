package main

import (
	"crypto/md5"
	"encoding/hex"
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

//in memory db

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

func main() {
	fmt.Println("we are creating url shortener ")
	OriginalUrl := "https://github.com/palindrome-kasak"
	genrateShortUrl(OriginalUrl)
	// setup server , local host
	// http server start
	fmt.Println("starting the http server on port 3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error on starting server:,", err)
	}
}
