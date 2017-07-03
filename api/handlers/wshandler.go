package handlers

import (
	"fmt"
	"time"

	"golang.org/x/net/websocket"
)

var Connect = ConnectionChan{
	removeConnection: make(chan *Client),
	addConnection:    make(chan *websocket.Conn),
	newMassege:       make(chan *Message),
	listConnect:      make(chan bool),
	count:            0,
}

func (connect *ConnectionChan) StartConnection() {

	for {
		select {
		case conn := <-connect.addConnection:

			AllConnection.ActiveCLient[conn] = LoginClient
			connect.count++

			fmt.Println(connect.count)

			for k := range AllConnection.ActiveCLient {
				websocket.JSON.Send(k, ListActiveCliens(AllConnection))
			}
			defer close(connect.addConnection)
		case conn := <-connect.removeConnection:
			for k, v := range AllConnection.ActiveCLient {
				if conn.Inner.LoginName == v.Inner.LoginName {

					delete(AllConnection.ActiveCLient, k)
					k.Close()
				}
				connect.count--

			}
			for k := range AllConnection.ActiveCLient {
				websocket.JSON.Send(k, ListActiveCliens(AllConnection))
			}
			defer close(connect.removeConnection)
		case message := <-connect.newMassege:
			for k, v := range AllConnection.ActiveCLient {
				if v.Inner.LoginName == message.Inner.ToLoginName || v.Inner.LoginName == message.Inner.FromLoginName {
					sendMessage(k, message)
					MessageSaveToClient(message, v)
				}

			}
			defer close(connect.newMassege)
		}

	}
}
func HandleChatRoom(ws *websocket.Conn) {
	//var newMessageFromWS Message

	var temp *websocket.Conn
	var newMessageFromWs Message
	temp = ws
	Connect.addConnection <- temp
	fmt.Println("Login user #6")
	for {

		if err := websocket.JSON.Receive(ws, &newMessageFromWs); err == nil {
			fmt.Println(newMessageFromWs.Inner.ToLoginName)
			newMessageFromWs.Inner.FromLoginName = AllConnection.ActiveCLient[ws].Inner.LoginName
			newMessageFromWs.Inner.TimeStamp = time.Now()
			Connect.newMassege <- &newMessageFromWs
			time.Sleep(time.Second * 2)
		}

	}
}
func sendMessage(ws *websocket.Conn, message *Message) {
	websocket.JSON.Send(ws, message)
}
