package main

import (
	"fmt"
	"io/fs"
	"net/http"

	"embed"
)

//go:embed static/*
var static embed.FS

const PORT = 3000

func main() {
	staticSub, _ := fs.Sub(static, "static")
	staticFs := http.FS(staticSub)

	http.Handle("/", http.FileServer(staticFs))

	http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
}
