package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)


/**
 * request:
 * | - header size - | - header len - | .. | key size | key    | value len | .. | value size | value  | .. | .. |
 * |     5 byte      |     5 byte     |    |  4 byte  | N byte | 1 byte    |    | 1 byte     | N byte |    |    |
 *
 * ===============================================================================================================
 */
// TCP 客户端
func main() {

	conn, err := net.Dial("tcp4", "127.0.0.1:9999")
	if err != nil {
		fmt.Println("err : ", err)
		return
	}
	defer conn.Close() // 关闭TCP连接
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n') // 读取用户输入
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" { // 如果输入q就退出
			return
		}
		_, err := conn.Write([]byte(inputInfo)) // 发送数据
		if err != nil {
			return
		}
		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("recv failed, err:", err)
			return
		}
		fmt.Println("header",string(buf[:5]),"\n")
		fmt.Println("header key",string(buf[:4]),"\n")
		fmt.Println("value",string(buf[4:5]),"\n")
		fmt.Println(string(buf[:n]))
	}
}
