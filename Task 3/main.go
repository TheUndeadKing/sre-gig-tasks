package main

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	p := "./assets/index.html"
	http.ServeFile(w, r, p)
}

func main() {

	http.HandleFunc("/", Home)

	http.ListenAndServe(":8080", nil)

}
