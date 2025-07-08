package main

import (
	"encoding/json"
	"fmt"
	"geerpc"
	"geerpc/codec"
	"log"
	"net"
	"time"
)

// 启动服务端，监听端口
func startServer(addr chan string) {
	// pick a free port
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal("network error:", err)
	}
	log.Println("start rpc server on", l.Addr())
	addr <- l.Addr().String()
	geerpc.Accept(l)
}

func main() {
	// 1. 启动服务器
	log.SetFlags(0)
	addr := make(chan string)
	go startServer(addr)

	// 2. 启动客户端发送连接
	// in fact, following code is like a simple geerpc client
	// 会开启多个连接，交给accept处理。accept是一个react模型，会先采用纤程的方式异步处理请求
	conn, _ := net.Dial("tcp", <-addr)
	defer conn.Close()

	//time.Sleep(time.Second)
	// 3. 首先发送Options，使用默认的GobType进行编码
	_ = json.NewEncoder(conn).Encode(geerpc.DefaultOption)
	cc := codec.NewGobCodec(conn)

	// 4. 发送消息并接受响应
	for i := 0; i < 5; i++ {
		// 4.1 发送消息头header和消息体 geerpc req
		h := &codec.Header{
			ServiceMethod: "Foo.Sum",
			Seq:           uint64(i),
		}
		_ = cc.Write(h, fmt.Sprintf("geerpc req %d", h.Seq))

		// 4.2 接收消息头和消息体
		_ = cc.ReadHeader(h)
		var reply string
		_ = cc.ReadBody(&reply)
		log.Println("reply:", reply)
	}
	time.Sleep(time.Second * 5)
}
