package main

import (
	"fmt"
	"go-csrf/pkg/domain"
	"log"
	"net/http"
)

func main() {

	csrf := domain.GenRandSlug(32)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, csrf)
		return
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Print(err)
	}

	fmt.Println("finished")
}
