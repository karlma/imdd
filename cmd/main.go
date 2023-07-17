package main

import (
	"fmt"
	"imd/internal/handler"
	"net/http"
)

func main() {
	fmt.Printf("start imd\n")
	http.HandleFunc("/send", handler.SendHandler)
	http.ListenAndServe(":8080", nil)
}
