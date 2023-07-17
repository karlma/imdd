package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fmt.Printf("start imd\n")
	http.HandleFunc("/send", sendHandler)
	http.ListenAndServe(":8080", nil)

}

func sendHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		fmt.Fprint(w, "The method is POST\n")
		return
	}
	defer req.Body.Close()

	b, _ := io.ReadAll(req.Body)

	f, err := os.Create("./data/msg.txt") // TODO: 自动创建目录的功能
	if err != nil {
		fmt.Fprintf(w, "Open file error: %v\n", err)
		return
	}

	n, err := f.Write(b)
	if err != nil {
		fmt.Fprintf(w, "Write file error: %v\n", err)
		return
	}
	fmt.Printf("wrote %d bytes\n", n)
	f.Sync()

	fmt.Fprintf(w, "OK")
}
