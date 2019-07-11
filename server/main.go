package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", handlePayload)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}

func handlePayload(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var b bytes.Buffer
	if _, err := io.Copy(&b, r.Body); err != nil {
		fmt.Printf("[ERR] failed to copy body: %s\n", err)
	}

	fmt.Println(b.String())
}
