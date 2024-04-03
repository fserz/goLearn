package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int
	// 在线用户列表
	OnLineMap map[string]*User
	mapLock   sync.RWMutex
	// 消息广播的channel
	Message chan string
}

// 创建一个server接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnLineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

// 监听Message广播消息channel的goroutine，一旦有消息就全部发送给在线User
func (this *Server) ListenMessager() {
	for {
		msg := <-this.Message

		// 将msg发送给全部在线的User
		this.mapLock.Lock()
		for _, cli := range this.OnLineMap {
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}
}

func (this *Server) Handler(conn net.Conn) {
	//当前连接业务
	//fmt.Println("连接建立成功")
	user := NewUser(conn, this)

	user.Online()

	// 监听用户是否活跃的channel
	isLive := make(chan bool)

	// 接受客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}
			// 有错误，每一次末尾是EOF
			if err != nil && err != io.EOF {
				fmt.Println("Conn Read err: ", err)
				return
			}

			// 提取用户的消息(去除'\n')
			msg := string(buf[:n-1])

			// 用户将消息进行处理
			user.DoMessage(msg)

			// 用户的任意消息，代表用户是一个活跃的
			isLive <- true
		}
	}()
	// 当前handler阻塞

	for {
		select {
		case <-isLive:
			//当前用户活跃，重置定时器

		case <-time.After(time.Second * 300):
			// 已经超时
			// 将User强制关闭

			user.sendMsg("u r out")

			// 销毁用的资源
			close(user.C)

			// 关闭连接
			conn.Close()

			// 退出当前的Handler
			return

		}
	}
}

// 广播消息的方法
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "] " + user.Name + msg
	this.Message <- sendMsg
}

// 启动服务器的接口
func (this *Server) Start() {
	// socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err :", err)
		return
	}

	// close listen socket
	defer listener.Close()

	// 启动监听的Message的goroutine
	go this.ListenMessager()

	for {
		// accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err:", err)
			continue
		}

		// do handler
		go this.Handler(conn)
	}
}
