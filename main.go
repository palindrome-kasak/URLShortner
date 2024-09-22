package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

type URL struct {
	Id           int       `json:"Id"`
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

func main() {
	fmt.Println("we are creating url shortener ")
	OriginalUrl := "https://github.com/palindrome-kasak"
	genrateShortUrl(OriginalUrl)
}
