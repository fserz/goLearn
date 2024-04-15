package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int //当前client的模式
}

func NewClient(serverIp string, serverPort int) *Client {
	// 创建客户端
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flag:       999,
	}

	// 创建server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial error: ", err)
		return nil
	}
	client.conn = conn
	// 返回对象
	return client
}

// 处理server回应的消息，直接显示到标准输出即可
func (client *Client) DealResponse() {
	// 一旦client.conn有数据，就直接copy到stdout标准输出上
	io.Copy(os.Stdout, client.conn)

	//上面的代码等价于这个
	//for {
	//	buf := make([]byte, 4096)
	//	client.conn.Read(buf)
	//	fmt.Println(buf)
	//}
}

func (client *Client) menu() bool {
	var flag int

	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.update username")
	fmt.Println("0.exit")

	fmt.Scanln(&flag)
	if flag >= 0 && flag <= 3 {
		client.flag = flag
		return true
	} else {
		fmt.Println(">>>>>>>>>>>>请输入合法范围数字<<<<<<<<<<<")
		return false
	}
}

// 查询在线用户
func (client *Client) SelectUsers() {
	sendMsg := "who\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn Write err: ", err)
		return
	}
}

// 私聊模式
func (client *Client) PrivateChat() {
	var remoteName string
	var chatMsg string

	client.SelectUsers()
	fmt.Println("请输入聊天聊天对象[用户名]，exit退出: ")
	fmt.Scanln(&remoteName)

	for remoteName != "exit" {
		fmt.Println(">>>>>>>请输入消息内容，exit退出: ")
		fmt.Scanln(&chatMsg)

		for chatMsg != "exit" {
			if len(chatMsg) != 0 {
				sendMsg := "to|" + remoteName + "|" + chatMsg + "\n\n"
				_, err := client.conn.Write([]byte(sendMsg))
				if err != nil {
					fmt.Println("conn Write err：", err)
					break
				}
			}
			chatMsg = ""
			fmt.Println(">>>>>>>请输入消息内容，exit退出: ")
			fmt.Scanln(&chatMsg)
		}

		client.SelectUsers()
		fmt.Println("请输入聊天聊天对象[用户名]，exit退出: ")
		fmt.Scanln(&remoteName)
	}
}

func (client *Client) PublicChat() {
	// 提示用户输入消息
	var chatMsg string

	fmt.Println(">>>>>>>>>>请输入聊天内容，exit退出<<<<<<<<<<")
	// &chatMsg 表示将用户输入的内容存储到 chatMsg 变量的地址中，
	// 这样可以在函数内部修改 chatMsg 变量的值。
	fmt.Scanln(&chatMsg)

	for chatMsg != "exit" {
		// 发给服务器

		// 消息不空则发送
		if len(chatMsg) != 0 {
			sendMsg := chatMsg + "\n"
			_, err := client.conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("conn Write err：", err)
				break
			}
		}
		chatMsg = ""
		fmt.Println(">>>>>>>>>>>>>请输入聊天的内容<<<<<<<<<<<<<, exit退出")
		fmt.Scanln(&chatMsg)
	}

}

func (client *Client) UpdateName() bool {
	fmt.Println(">>>>>>>>>>>>>>>>请输入用户名<<<<<<<<<<<<<<<")
	fmt.Scanln(&client.Name)
	sendMsg := "rename|" + client.Name + "\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write err: ", err)
		return false
	}
	return true
}

func (client *Client) Run() {
	for client.flag != 0 {
		for client.menu() != true {
		}
		// 根据不同的模式处理不同的业务
		switch client.flag {
		case 1:
			// 公聊模式
			client.PublicChat()
			break
		case 2:
			// 私聊模式选择
			client.PrivateChat()
			break
		case 3:
			// 更新用户名
			client.UpdateName()
			break
		}

	}
}

var serverIp string
var serverPort int

// ./client -ip 127.0.0.1 -port 8888
func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "Set the server IP address(default  127.0.0.1)")
	flag.IntVar(&serverPort, "port", 8888, "Set the server port(default  8888)")
}

func main() {

	// 命令行解析
	flag.Parse()

	client := NewClient(serverIp, serverPort)
	//client := NewClient("127.0.0.1", 8888)

	if client == nil {
		fmt.Println(">>>>>>>>>>>>Connect Server Failed<<<<<<<<<<")
		return
	}

	// 单独开启一个goroutine去处理server的回执消息
	go client.DealResponse()

	fmt.Println(">>>>>>>>>>>>Connect Server Successfully<<<<<<<<<<")

	// 启动客户端的业务
	client.Run()

}
