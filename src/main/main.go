package main

import "net/http"

type person struct {
	name string
}

func (p *person) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("First Name " + p.name))
}

func main() {
	personOne := &person{name: "Julio"}
	http.ListenAndServe(":8080", personOne)
}