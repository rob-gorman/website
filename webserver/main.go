package webserver

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	PORT := 3000
	s := http.NewServeMux()
	s.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "We are in our WebServer Yes?\n")
	})

	http.ListenAndServe(fmt.Sprintf(":%d", PORT), s)
}
