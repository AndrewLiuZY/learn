package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
)

func main() {
	md5 := md5.New()
	data := []byte("Hello im lzy")
	n, err := md5.Write(data)
	if n != len(data) || err != nil {
		log.Printf("Hash write error: %v / %v", n, err)
	}
	b := []byte{}
	checksum := md5.Sum(b)
	fmt.Printf("Result: %x\n", checksum)
	md5.Reset()
	io.WriteString(md5, "Hello im lzy")
	str := md5.Sum(b)
	fmt.Println(str)
	fmt.Printf("%x", str)
}
