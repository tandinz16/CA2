package main

import (
	"encoding/json"
	"log"
	"net/http"
	"rockpaperscs/rps"
	"strconv"
	"text/template"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Println(err)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}
}

func playHandler(w http.ResponseWriter, r *http.Request) {
	// extract the user choice value from query string
	str_value := r.URL.Query().Get("c")
	// convert string to int
	userChoice, _ := strconv.Atoi(str_value)
	result := rps.PlayRound(userChoice)
	jsObj, err := json.Marshal(result)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(jsObj)
}
func main() {
	var port = 8080
	http.HandleFunc("/", homePage)
	http.HandleFunc("/play", playHandler)

	log.Println("Starting web server on Port ", port)
	http.ListenAndServe(":8080", nil)
}
