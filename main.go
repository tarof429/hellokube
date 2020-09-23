package main

import (
	"fmt"
	"net/http"
	"os"
)

func myhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Go!\n")
}

func main() {
	http.HandleFunc("/", myhandler)

	err := http.ListenAndServe(":8090", nil)

	if err != nil {
		fmt.Printf("Unable to start serevr")
		os.Exit(1)
	}
}
