package node

import (
	"bufio"
	"fmt"
	"hb_process/utils"
	"net"
	"os"
	"path"
	"syscall"
)

type AppAgent struct {
	listen  net.Listener
	unix    string
	appConn net.Conn
}

func (a *AppAgent) GetAddress() string {
	return a.unix
}

func (a *AppAgent) SetUnix(unix string) {
	a.unix = unix
}

func NewAppAgent(tag string) *AppAgent {
	connectId := utils.GetRandomString(16)
	if !utils.FileExists(path.Join("./", "app_nodes")) {
		_ = os.MkdirAll(path.Join("./", "app_nodes"), 0755)
	}
	return &AppAgent{unix: path.Join("./", "app_nodes", fmt.Sprintf("%s-%s.sock", tag, connectId))}
}

func (a *AppAgent) ListenAndServe() (err error) {
	a.listen, err = net.Listen("unix", a.unix)

	fmt.Println(a.unix, err)
	if err != nil {
		return err
	}

	// 关闭时删除绑定的文件
	defer func() {
		a.Close()
	}()

	//c := make(chan net.Conn, 3)
	for {
		conn, err := a.listen.Accept()
		if err != nil {
			return err
		}

		go func() {
			//defer conn.Close()
			for {
				reader := bufio.NewReader(conn)
				var buf [4096]byte
				n, err := reader.Read(buf[:]) // 读取数据
				if err != nil {
					fmt.Println("read from client failed, err: ", err)
					break
				}
				recvStr := string(buf[:n])
				fmt.Println("收到UNIX端发来的数据：", recvStr)
				conn.Write([]byte("海宝")) // 发送数据

				// 简单实现双向通信
				// 保存conn 用于下一次unix通信的调用 接到unixClient 先保存  下次 在保存一次

				//c <- conn
				//if c != nil {
				//conn := <-c
				//for i:=0;i<=3;i++{
				//	conn.Write([]byte("海"))
				//}
				//}

				// 存入php子进程
				//childNode := &PHPAppNode{agent: a, pid: 1}
			}
		}()
	}
}

func (a *AppAgent) Close() {
	if a.appConn != nil {
		_ = a.appConn.Close()
	}
	_ = a.listen.Close()
	_ = syscall.Unlink(a.unix)
}
