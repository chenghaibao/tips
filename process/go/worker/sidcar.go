package worker

import (
	"bufio"
	"context"
	"fmt"
	"hb_process/worker/node"
	"net"
)

type Sidecar struct {
	ctx         context.Context
	cancel      context.CancelFunc
	listener    net.Listener // 对外tcp服务
	appNodeManage *node.PHPAppNodeManage
	initialized bool
}

func NewSidecar(parentCtx context.Context) *Sidecar {
	ctx, cancelFunc := context.WithCancel(parentCtx)
	return &Sidecar{
		ctx:    ctx,
		cancel: cancelFunc,
	}
}

func (s *Sidecar) Init() (err error) {
	listenAddr := fmt.Sprintf("%s:%d", "127.0.0.1", 9997)
	s.listener, err = net.Listen("tcp", listenAddr)


	if err != nil {
		return err
	}

	err = s.startAppNodeManage()
	if err != nil {
		return err
	}

	s.startWorkMange()
	return nil
}

func (s *Sidecar) Server() error {
	for {
		conn, err := s.listener.Accept() // 监听客户端的连接请求
		if err != nil {
			return err
		}
		// 启动一个goroutine来处理客户端的连接请求
		go func(){
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
				fmt.Println("收到TCP端发来的数据：", recvStr)
				conn.Write([]byte("海宝")) // 发送数据
			}
		}()
	}
}

func (s *Sidecar) Close() {
	s.listener.Close()
}

// 开启phpMaster worker
func (s *Sidecar) startAppNodeManage() (err error) {
	node.NewAppNodeManage()
	return nil
}


func (s *Sidecar) startWorkMange(){

}