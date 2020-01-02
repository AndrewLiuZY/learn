package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		res, err := http.Get("http://" + url)
		checkError(err)
		data, err := ioutil.ReadAll(res.Body)
		checkError(err)
		fmt.Printf("Got: %q", string(data))
	}

}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Get : %v", err)
	}
}
