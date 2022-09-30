package main

import (
	"fmt"
	"log"
	"net/http"
)

func formhandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "parseform() err: %v", err)
		return
	}
	fmt.Fprintf(w, "post request succes")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "name=%s\n", name)
	fmt.Fprintf(w, "address=%s\n", address)
}
func hellohandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not founds", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "not founds", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello")
}

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formhandler)
	http.HandleFunc("/hello", hellohandler)
	fmt.Println("starting at 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
