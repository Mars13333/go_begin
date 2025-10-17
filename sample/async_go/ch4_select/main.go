package main

import (
	"fmt"
	"time"
)

type MessageServer struct {
	mesageCh chan string
	quitCh   chan struct{}
}

func NewMessageServer() *MessageServer {
	return &MessageServer{
		mesageCh: make(chan string, 100), // 业务channel尽量设置
		quitCh:   make(chan struct{}),
	}
}

// 方法 处理消息
func (s *MessageServer) handleMessage(msg string) {
	fmt.Println("Handling message:", msg)
}

// 函数 写入消息
func sendMessage(s *MessageServer, message string, number int) {
	for i := 0; i < number; i++ {
		s.mesageCh <- fmt.Sprintf("message %d: %s", i+1, message)
	}
}

// 方法 loop死循环 工作循环
func (s *MessageServer) work() {
messageLoop:
	for {
		select {
		case message := <-s.mesageCh:
			s.handleMessage(message)
		case <-s.quitCh:
			// 收到退出信号，退出
			fmt.Println("quitting server...")
			break messageLoop
			// default:
			// 没有消息
			// 可以sleep
			// 可以做其他事情，比如心跳检测等
			// 可以用ticker定时器来做周期性任务

		}
	}
	fmt.Println("server is down...")
}

func (s *MessageServer) quit() {
	s.quitCh <- struct{}{} // 0 byte, 经常作为结束信号
}

func main() {
	server := NewMessageServer()
	go func() {
		time.Sleep(time.Second * 3)
		server.quit()
	}()
	sendMessage(server, "Hello, World!", 5)
	server.work()
}
