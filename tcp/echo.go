package tcp

import (
	"bufio"
	"context"
	"go-redis/lib/logger"
	"go-redis/lib/sync/atomic"
	"go-redis/lib/sync/wait"
	"io"
	"net"
	"sync"
	"time"
)

/**
测试服务是否正常运行
*/

// EchoHandler 维护了一个client的map，还有是否关闭closing字段
type EchoHandler struct {
	activeConn sync.Map
	closing    atomic.Boolean
}

func MakeHandler() *EchoHandler {
	return &EchoHandler{}
}

func (h *EchoHandler) Handler(ctx context.Context, conn net.Conn) {
	// 如果是链接断开
	if h.closing.Get() {
		_ = conn.Close()
	}
	client := &EchoClient{
		Conn: conn,
	}
	// 设置键的值
	h.activeConn.Store(client, struct{}{})
	reader := bufio.NewReader(conn)

	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				logger.Info("connect close")
				h.activeConn.Delete(client)
			} else {
				logger.Warn(err)
			}
			return
		}
		// 客户端的waitGroup +1
		client.Waiting.Add(1)
		b := []byte(msg)
		_, _ = conn.Write(b)
		client.Waiting.Done()
	}

}

func (h *EchoHandler) Close() error {
	logger.Info("handler shutting down ...")
	h.closing.Set(true)
	h.activeConn.Range(func(key interface{}, value interface{}) bool {
		client := key.(*EchoClient)
		_ = client.Close()
		return true

	})
	return nil
}

// EchoClient 一个连接客户端，
type EchoClient struct {
	Conn    net.Conn
	Waiting wait.Wait
}

// Close connection  设置连接的等待时间
func (c *EchoClient) Close() error {
	c.Waiting.WaitWithTimeout(10 * time.Second)
	_ = c.Conn.Close()
	return nil
}
