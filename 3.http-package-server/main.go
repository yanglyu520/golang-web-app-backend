package main

import (
	"fmt"
	//"io/ioutil"
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", handler{})
}

type handler struct{}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	fmt.Println(r.URL.Path)

	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(`{"hello": "dave"}`))



		// b, _ := ioutil.ReadAll(r.Body)
	// fmt.Println(string(b))

	// get it from the database
}