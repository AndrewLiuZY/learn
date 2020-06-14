package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	//打开连接:
	conn, err := net.Dial("tcp", "localhost:50000")
	checkErr(err, "Error dialing")
	buf := make([]byte, 512)
	l, err := conn.Read(buf)
	checkErr(err, "Error reading")
	if l != 0 {
		recieveStr := string(buf[:l])
		fmt.Println(recieveStr)
	}
	//由于目标计算机积极拒绝而无法创建连接
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("First, what is your name?")
	clientName, _ := inputReader.ReadString('\n')
	// fmt.Printf("CLIENTNAME %s", clientName)
	trimmedClient := strings.Trim(clientName, "\r\n") // Windows 平台下用 "\r\n"，Linux平台下使用 "\n"
	go func() {
		for {
			buf := make([]byte, 512)
			l, err := conn.Read(buf)
			checkErr(err, "Error reading")
			if l != 0 {
				recieveStr := string(buf[:l])
				fmt.Println(recieveStr)
			}
		}
	}()
	// 给服务器发送信息直到程序退出：
	for {
		fmt.Println("What to send to the server? Type Q to quit.")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")
		// fmt.Printf("input:--s%--", input)
		// fmt.Printf("trimmedInput:--s%--", trimmedInput)~
		if trimmedInput == "Q" {
			return
		}
		if trimmedInput == "WHO" {
			_, err = conn.Write([]byte("WHO"))
		} else {
			_, err = conn.Write([]byte(trimmedClient + " says: " + trimmedInput))
		}
	}
}
func recieveData() {

}

func checkErr(err error, notice string) {
	if err != nil {
		fmt.Println(notice + err.Error())
		os.Exit(2)
	}
}

var methodChannel chan func() = make(chan func(), 2)

func excute(f func()) {
	methodChannel <- f
}

func init() {
	go func() {
		for {
			(<-methodChannel)()
		}
	}()
}
