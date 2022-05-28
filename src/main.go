package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.String())
	msg := "<html><body>Hello World!</body></html>"
	w.Write([]byte(msg))
}

func main() {
	http.HandleFunc("/", hello)
	fmt.Println("listen :3000")
	http.ListenAndServe(":3000", nil)
}
