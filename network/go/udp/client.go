package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("udp", ":8000")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer conn.Close()

	go func() {
		str := make([]byte, 2048)
		n, _ := os.Stdin.Read(str)
		_, err1 := conn.Write(str[:n])
		if err1 != nil {
			fmt.Println(err1.Error())
		}
	}()
	buf := make([]byte, 2048)
	for {
		//读取数据
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("读到的数据：" + string(buf[:n]))
	}

}