package main

import (
	"fmt"
	nameWriter "go-server/NameWriter"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hey bitch - route: ")
	})

	http.HandleFunc("/art/{name}", func(w http.ResponseWriter, r *http.Request) {
		art_name := r.PathValue("name")
		http.ServeFile(w, r, "assets/"+art_name+".txt")
	})

	http.HandleFunc("/write-my-name/{name}", func(w http.ResponseWriter, r *http.Request) {
		art_name := r.PathValue("name")
		style := r.URL.Query().Get("style")
		color := r.URL.Query().Get("color")

		res := nameWriter.WriteName(art_name, style, color)
		fmt.Println(res)
		fmt.Fprintf(w, res)
	})

	http.HandleFunc("/get-fonts", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, nameWriter.GetFonts())
	})

	http.ListenAndServe(":8000", nil)
}
