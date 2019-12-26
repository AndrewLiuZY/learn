package main

import "fmt"

type Request struct {
	a, b   int
	replyc chan int // 请求中的回复频道

}

type binOp func(a, b int) int

func run(op binOp, req *Request) {

	req.replyc <- op(req.a, req.b)

}

func server(op binOp, service chan *Request) {

	for {

		req := <-service // 请求到达这里

		// 开启请求的 Goroutine ：

		go run(op, req) // 不要等待 op

	}

}

func startServer(op binOp) chan *Request {

	reqChan := make(chan *Request)

	go server(op, reqChan)

	return reqChan

}

func main() {

	adder := startServer(func(a, b int) int { return a + b })

	const N = 10000

	var reqs [N]Request

	for i := 0; i < N; i++ {

		req := &reqs[i]

		req.a = i

		req.b = i + N

		req.replyc = make(chan int)

		adder <- req

	}

	// 校验：

	for i := N - 1; i >= 0; i-- { // 顺序无所谓

		if <-reqs[i].replyc != N+2*i {

			fmt.Println("fail at", i)

		} else {

			fmt.Println("Request ", i, "is ok!")

		}

	}

	fmt.Println("done")

}
