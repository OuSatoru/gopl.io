//use omdbapi.com api to get poster
//type name download poster
//http://www.omdbapi.com/?t={{title}}&plot=full& => Plot and Poster

package main

import (
	"net/url"
	"strings"
	"net/http"
	"fmt"
	"encoding/json"
	"os"
	"log"
	"io"
)

const OmdbUrl = "http://www.omdbapi.com/"

type MovieResult struct {
	Title    string
	Director string
	Plot     string
	Poster   string
}

func main() {
	result, err := searchTitle(strings.Join(os.Args[1:], " "))
	if err != nil {
		log.Fatal(err)
	}
	poster, err := os.Create(result.Title + ".jpg")
	if err != nil {
		panic(err)
	}
	defer poster.Close()
	res, err := http.Get(result.Poster)
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(poster, res.Body)
	fmt.Println(result.Title + ": " + result.Plot)
}

func searchTitle(title string) (*MovieResult, error) {
	//title = strings.Replace(title, " ", "+", -1) no need
	title = url.QueryEscape(title)
	res, err := http.Get(OmdbUrl+"?t="+title+"&plot=full&")
	//res, err := http.Get(`http://www.omdbapi.com/?t=kimi+no+na+wa&plot=full&`)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", res.Status)
	}
	var movie MovieResult
	if err := json.NewDecoder(res.Body).Decode(&movie); err!=nil {
		return nil, err
	}
	return &movie, nil
}