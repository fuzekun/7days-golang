package main

import (
	"fmt"
	"geerpc"
	"log"
	"net"
	"sync"
	"time"
)

/*
*
addr : 协程之间传递参数
*/
func startServer(addr chan string) {
	// pick a free port
	// 1. 监听通道，随机选择一个可用的port
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal("network error:", err)
	}
	log.Println("start rpc server on", l.Addr())
	// 2. 返回开启的port
	addr <- l.Addr().String()
	// 3. 进行监听
	geerpc.Accept(l)
}

func main() {
	// 1. 连接服务器
	log.SetFlags(0)
	// 1.1 绑定通道
	addr := make(chan string)
	// 1.2 开启服务器
	go startServer(addr)
	// 1.3 与rpc服务器建立tcp连接
	client, _ := geerpc.Dial("tcp", <-addr)
	defer func() { _ = client.Close() }()

	time.Sleep(time.Second)
	// send request & receive response
	// 2. 发送请求，接收响应
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			args := fmt.Sprintf("geerpc req %d", i)
			var reply string
			if err := client.Call("Foo.Sum", args, &reply); err != nil {
				log.Fatal("call Foo.Sum error:", err)
			}
			log.Println("reply:", reply)
		}(i)
	}
	wg.Wait()
}
