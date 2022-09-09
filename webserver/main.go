package main

import (
	"fmt"
	"io/fs"
	"net/http"

	"embed"
)

// content holds our static web server content.
// //go:embed static/index.html
// var index []byte

// //go:embed static/images/*
// var images embed.FS

// //go:embed static/assets/*
// var assets embed.FS

//go:embed static/*
var static embed.FS

func main() {
	PORT := 3000
	// assetsFs := http.FileServer(http.FS(assets))
	// imagesFs := http.FileServer(http.FS(images))

	staticSub, _ := fs.Sub(static, "static")
	staticFs := http.FS(staticSub)
	// s := mux.NewRouter()

	http.Handle("/", http.FileServer(staticFs))

	// s.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.WriteHeader(http.StatusOK)
	// 	w.Write(index)
	// })
	// s.PathPrefix("/static/assets/").Handler(http.StripPrefix("/static/", assetsFs))
	// s.PathPrefix("/static/images/").Handler(http.StripPrefix("/static/", imagesFs))

	// s.Handle("/static/", http.FileServer(staticContent)).Methods("GET")
	// s.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static")))).Methods("GET")
	// s.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "./static/index.html")
	// }).Methods("GET")

	// s.Handle("/", http.FileServer(staticContent))

	http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
}
