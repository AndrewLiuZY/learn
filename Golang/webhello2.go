package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.Handle("/hello/", http.HandlerFunc(helloServer))
	http.HandleFunc("/shouthello/", shoutHelloServer)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err.Error())
	}
}

func helloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside helloServer handler")
	path := req.URL.Path
	fmt.Fprintf(w, "hello "+path[strings.LastIndex(path, "/")+1:])
}

func shoutHelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside helloServer handler")
	path := req.URL.Path
	fmt.Fprintf(w, "hello "+strings.ToUpper(path[strings.LastIndex(path, "/")+1:]))
}
