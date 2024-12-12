package main

import (
	"io"
	"log"
	"net/http"
)

func feature1(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Feature5")
}

func feature2(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Feature4")
}
func main() {
	//http.HandleFunc("/feature1", feature1)
	//http.HandleFunc("/feature2", feature2)
	//http.ListenAndServe(":8080", nil)

	err := http.ListenAndServe(":8080", handlers())
	if err != nil {
		log.Fatal(err)
	}
}

func handlers() http.Handler {
	h := http.NewServeMux()
	h.HandleFunc("/feature5", feature1)
	h.HandleFunc("/feature4", feature2)
	return h
}
