package main

import (
	"fmt"
	"net"
)

func main() {
	//创建udp地址
	udp_addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err.Error())
	}
	//监听udp地址
	udp_conn, err := net.ListenUDP("udp", udp_addr)
	if err != nil {
		fmt.Println(err.Error())
	}
	//延迟关闭监听
	defer udp_conn.Close()

	for {
		buf := make([]byte, 1024)
		//阻塞获取数据
		n, rAddr, err := udp_conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("客户端【%s】,发送的数据为：%s\n", rAddr, string(buf[:n]))
		//发送给客户端数据
		_, errs := udp_conn.WriteToUDP([]byte("nice to meet you"), rAddr)
		if errs != nil {
			fmt.Println(err.Error())
		}
	}
}