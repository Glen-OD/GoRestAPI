package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Article struct {
	X      string `json:"x"`
	Answer int    `json:"answer"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		//Article{Title: "Test Title", Desc: "Test Description", Content: "Hello World"},
		//Article{answer: "Hello World"},
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
	x := keys[0]

	//tester
	//log.Println("the key is " + key)
	//log.Println(strings.Count(key, ","))
	//answer := strconv.Itoa(strings.Count(x, ","))
	answer := strings.Count(x, ",")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	//Data := []byte(`[
	//    {"Name": key},
	//    {"answer": numberOfCommas}
	// ]`)

	//	user := &Article{x: x,
	//		answer: answer}
	//fmt.Fprintf(w, "The number of commas is: "+string(numberOfCommas))
	//fmt.Fprintf(w, "The sentence is: "+string(key))

	//fmt.Fprintf(w, "Homepage Endpoint Hit")

	foo_marshalled, err := json.Marshal(Article{X: x, Answer: answer})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprint(w, string(foo_marshalled))

	//json.NewEncoder(w).Encode(user)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/test", allArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
