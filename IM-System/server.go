package main

import (
	"fmt"
	"net"
	"sync"
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
	//当前链接业务
	//fmt.Println("连接建立成功")
	user := NewUser(conn)
	// 用户上线，加入到OnlineMap
	this.mapLock.Lock()
	this.OnLineMap[user.Name] = user
	this.mapLock.Unlock()

	// 广播用户上线消息
	this.BroadCast(user, "已上线")

	// 当前handler阻塞
	select {}
}

// 广播消息的方法
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + msg
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
