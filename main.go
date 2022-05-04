package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	redirectTo := os.Getenv("REDIRECT_TO")
	fmt.Println("REDIRECT_TO:", redirectTo)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, redirectTo, http.StatusTemporaryRedirect)
	})
	err := http.ListenAndServe("127.0.0.1:80", nil)
	if err != nil {
		panic(err)
	}
}
