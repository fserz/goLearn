package main

import (
	"net"
	"strings"
)

type User struct {
	Name   string
	Addr   string
	C      chan string ``
	conn   net.Conn
	server *Server
}

// 创建一个用户的API
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}

	go user.ListenMessage()

	return user
}

// 监听当前User channel的方法，一旦有消息发送给对端客户端
func (this *User) ListenMessage() {
	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + "\n"))
	}
}

// 用户的上线业务
func (this *User) Online() {

	// 用户上线，加入到OnlineMap
	this.server.mapLock.Lock()
	this.server.OnLineMap[this.Name] = this
	this.server.mapLock.Unlock()

	// 广播用户上线消息
	//this.server.BroadCast(this, "user online")
	this.DoMessage(" user online")
}

// 用户的下线业务
func (this *User) Offline() {

	// 用户下线，从OnlineMap删除
	this.server.mapLock.Lock()
	delete(this.server.OnLineMap, this.Name)
	this.server.mapLock.Unlock()

	// 广播用户下线消息
	//this.server.BroadCast(this, "user offline")
	this.DoMessage(" user offline")

}

// 给当前user对应的客户端发消息
func (this *User) sendMsg(msg string) {
	this.conn.Write([]byte(msg))
}

// 用户处理消息的业务
func (this *User) DoMessage(msg string) {
	if msg == "who" {
		// 查询当前在线用户有哪些
		this.server.mapLock.Lock()
		for _, user := range this.server.OnLineMap {
			onlineMsg := "[" + user.Addr + "]" + user.Name + " is online...\n"
			this.sendMsg(onlineMsg)
		}
		this.server.mapLock.Unlock()

	} else if len(msg) > 7 && msg[:7] == "rename|" {
		// 消息格式：rename|张三
		newName := strings.Split(msg, "|")[1]

		// 判断name是否存在
		_, ok := this.server.OnLineMap[newName]
		if ok {
			this.sendMsg("name is used\n")
		} else {
			this.server.mapLock.Lock()
			delete(this.server.OnLineMap, this.Name)
			this.server.OnLineMap[newName] = this
			this.server.mapLock.Unlock()

			this.Name = newName
			this.sendMsg("user name has been updated: " + this.Name + "\n")
		}
	} else if len(msg) > 4 && msg[:3] == "to|" {
		// 消息格式：to|jackie|消息内容
		// 1.获取对方用户名
		remoteName := strings.Split(msg, "|")[1]
		if remoteName == "" {
			this.sendMsg("The message format is incorrect, please use \"to|Jackie|Hello~ \" format \n")
		}

		// 2.根据用户名，得到对方User对象
		remoteUser, ok := this.server.OnLineMap[remoteName]
		if !ok {
			this.sendMsg("User name does not exist\n")
			return
		}
		// 3.获取消息内容，通过对方的User对象将消息内容发送过去
		content := strings.Split(msg, "|")[2]
		if content == "" {
			this.sendMsg("no content, please send again")
		}
		remoteUser.sendMsg(this.Name + " say to u: " + content)
	} else {
		this.server.BroadCast(this, msg)
	}
}
