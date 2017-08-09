package main

import (
	"net/http"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://nytimes.com", 301)
}

func main() {
	http.HandleFunc("/", redirect)
	err := http.ListenAndServe(":90901", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
