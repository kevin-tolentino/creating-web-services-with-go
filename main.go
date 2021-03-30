package main

import (
	"fmt"
	"net/http"
)

type fooHandler struct {
	Message string
}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(f.Message))
	fmt.Println("a call was made")
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bar was called using barHandler 😀"))
}

func main() {
	http.Handle("/foo", &fooHandler{Message: "foo was called"})
	http.HandleFunc("/bar", barHandler)
	http.ListenAndServe(":5000", nil)

}
