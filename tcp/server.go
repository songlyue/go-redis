package tcp

import (
	"context"
	"fmt"
	"go-redis/interface/tcp"
	"go-redis/lib/logger"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// Config stores tcp server properties
type Config struct {
	Address string
}

// ListenAndServeWithSignal 监听有信号的服务，绑定端口号和处理请求，阻塞到有停止信号过来
func ListenAndServeWithSignal(cfg *Config, handler tcp.Handler) error {
	closeChan := make(chan struct{})
	sigCh := make(chan os.Signal)
	// 指定的信号通知给sigCh
	signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

	// 监听收到退出信号发送给closeChan
	go func() {
		sig := <-sigCh
		switch sig {
		case syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			closeChan <- struct{}{}
		}
	}()
	// 开启tcp服务
	listen, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		return err
	}

	logger.Info(fmt.Sprintf("bind:%s,start listening", cfg.Address))

	// 处理服务
	ListenAndServe(listen, handler, closeChan)
	return nil

}

func ListenAndServe(listen net.Listener, handler tcp.Handler, closeChan chan struct{}) {

	// 监听closeChan通道
	go func() {
		//  closeChan 在收到系统退出信息时有消息进来
		<-closeChan
		logger.Info("shutting down")
		_ = listen.Close()
		_ = handler.Close()
	}()

	ctx := context.Background()
	var waitDone sync.WaitGroup
	for {
		conn, err := listen.Accept()
		if err != nil {
			break
		}
		// 处理链接
		logger.Info(fmt.Sprintf("accept link:%s", conn.RemoteAddr()))
		waitDone.Add(1)
		go func() {
			defer func() {
				waitDone.Done()
			}()
			handler.Handler(ctx, conn)
		}()
	}
	waitDone.Wait()
}
