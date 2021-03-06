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

	var bHandle bool

	if !ok || len(keys[0]) < 1 {
		bHandle = true
		//log.Println("Url Param 'key' is missing")
	}

	if bHandle == true {
		var x string = ""
		answer := strings.Count(x, ",")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		user := &Article{X: x, Answer: answer}
		json.NewEncoder(w).Encode(user)
	}
	if bHandle == false {
		x := keys[0]
		answer := strings.Count(x, ",")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		foo_marshalled, err := json.Marshal(Article{X: x, Answer: answer})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Fprint(w, string(foo_marshalled))
	}

	// Query()["key"] will return an array of items,
	// we only want the single item.

	//tester
	//log.Println("the key is " + x)

	//log.Println(strings.Count(key, ","))
	//answer := strconv.Itoa(strings.Count(x, ","))

	//log.Println("the answer is " + strconv.Itoa(answer))

	//Data := []byte(`[
	//    {"Name": key},
	//    {"answer": numberOfCommas}
	// ]`)

	//	user := &Article{x: x,
	//		answer: answer}
	//fmt.Fprintf(w, "The number of commas is: "+string(numberOfCommas))
	//fmt.Fprintf(w, "The sentence is: "+string(key))

	//fmt.Fprintf(w, "Homepage Endpoint Hit")

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
