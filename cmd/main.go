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

	fn := "./data/msg.txt"
	// TODO: 自动创建目录
	f, err := os.OpenFile(fn, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Fprintf(w, "Open file error: %v\n", err)
		return
	}

	n, err := f.Write(append(b, '\n'))
	if err != nil {
		fmt.Fprintf(w, "Write file error: %v\n", err)
		return
	}
	fmt.Printf("wrote %d bytes\n", n)
	f.Sync()

	fmt.Fprintf(w, "OK")
}
