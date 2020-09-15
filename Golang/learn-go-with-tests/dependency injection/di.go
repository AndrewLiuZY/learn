package main

import (
	"fmt"
	"io"
	"net/http"
)

func main(){
	http.ListenAndServe(":5000",http.HandlerFunc(MyGreeterHandler))
}

func MyGreeterHandler(w http.ResponseWriter,r *http.Request){
	Greet(w,"word")
}

func Greet(buffer io.Writer,content string)  {
	fmt.Fprintf(buffer,"Hello %s",content)
}