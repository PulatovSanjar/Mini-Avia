package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintln(w, "Mini-Avia API is alive")
		if err != nil {
			return
		}
	})
	_ = http.ListenAndServe(":8080", nil)
}
