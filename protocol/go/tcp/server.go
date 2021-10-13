package main

import (
	"bufio"
	"fmt"
	"net"
)

// TCP Server端测试
// 处理函数
func process(conn net.Conn) {
	defer conn.Close() // 关闭连接
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Println("read from client failed, err: ", err)
			break
		}
		//recvStr := string(buf[:n])
		fmt.Println(n)
		fmt.Println("收到Client端发来的数据：", string(buf[:1]))
		conn.Write([]byte("type1")) // 发送数据
	}
}

func main() {
	listen, err := net.Listen("tcp4", "127.0.0.1:9999")
	if err != nil {
		fmt.Println("Listen() failed, err: ", err)
		return
	}
	for {
		conn, err := listen.Accept() // 监听客户端的连接请求
		if err != nil {
			fmt.Println("Accept() failed, err: ", err)
			continue
		}
		go process(conn) // 启动一个goroutine来处理客户端的连接请求
	}
}
