package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Article struct {
	//Title   string `json:"Title"`
	//Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		//Article{Title: "Test Title", Desc: "Test Description", Content: "Hello World"},
		Article{Content: "Hello World"},
	}

	fmt.Println("Endpoint HitL All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["key"]

	if !ok || len(keys[0]) < 1 {
		//log.Println("Url Param 'key' is missing")
		return
	}

	// Query()["key"] will return an array of items,
	// we only want the single item.
	key := keys[0]

	//tester
	//log.Println("the key is " + key)
	//log.Println(strings.Count(key, ","))
	numberOfCommas := strconv.Itoa(strings.Count(key, ","))
	fmt.Fprintf(w, "The number of commas is: "+string(numberOfCommas))
	//fmt.Fprintf(w, "The sentence is: "+string(key))

	//fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/test", allArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
