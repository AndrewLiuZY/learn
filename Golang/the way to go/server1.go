package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

var users []string = make([]string, 0)

func main() {
	fmt.Println("Starting the server ...")
	// 创建 listener
	listener, err := net.Listen("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return //终止程序
	}
	userCh := make(chan string)
	go func() {
		for {
			users = append(users, <-userCh)
		}
	}()
	// 监听并接受来自客户端的连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return // 终止程序
		}
		userCh <- conn.RemoteAddr().String()
		go doServerStuff1(conn)
	}
}

func doServerStuff1(conn net.Conn) {
	_, err := conn.Write([]byte("Congratulations! you connect success"))
	checkErr1(err, "send err", true)
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		recieveStr := string(buf[:len])
		if recieveStr == "WHO" {
			_, err := conn.Write([]byte(strings.Join(users, "\n")))
			checkErr1(err, "send err", true)
			fmt.Println(conn.RemoteAddr().String(), "checked the user list just now!")
		} else {
			fmt.Printf("Received data: %v \n", string(buf[:len]))
		}
	}
}

func checkErr1(err error, notice string, exit bool) {
	if err != nil {
		fmt.Println(notice + err.Error())
		if exit {
			os.Exit(2)
		}
	}
}
