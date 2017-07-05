package handlers

import (
	"fmt"
	"time"

	"github.com/davidborka/chatApp/api/model"

	"github.com/davidborka/chatApp/api/dbconnect"
	"golang.org/x/net/websocket"
)

var (
	AllConnection = model.Connection{make(map[*websocket.Conn]model.Client)}
)

type ConnectionChan struct {
	removeConnection chan *model.Client
	addConnection    chan *websocket.Conn
	newMassege       chan *model.Message
	listConnect      chan bool
	count            int
}

var Connect = ConnectionChan{
	removeConnection: make(chan *model.Client),
	addConnection:    make(chan *websocket.Conn),
	newMassege:       make(chan *model.Message),
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

		case conn := <-connect.removeConnection:
			for k, v := range AllConnection.ActiveCLient {
				if conn.LoginName == v.LoginName {

					delete(AllConnection.ActiveCLient, k)
					k.Close()
				}
				connect.count--

			}
			for k := range AllConnection.ActiveCLient {
				websocket.JSON.Send(k, ListActiveCliens(AllConnection))
			}

		case message := <-connect.newMassege:

			for k, v := range AllConnection.ActiveCLient {

				if v.LoginName == message.ToLoginName {
					sendMessage(k, message)
				} else if v.LoginName == message.FromLoginName {

					sendMessage(k, message)
				}
			}
			time.Sleep(time.Second * 2)

		}

	}
}
func HandleChatRoom(ws *websocket.Conn) {
	//var newMessageFromWS Message
	go Connect.StartConnection()
	time.Sleep(time.Second * 3)
	var temp *websocket.Conn
	var newMessageFromWs model.Message
	temp = ws
	fmt.Print("add connect")
	Connect.addConnection <- temp
	fmt.Println("Login user #6")
	for {

		if err := websocket.JSON.Receive(ws, &newMessageFromWs); err == nil {
			fmt.Println(newMessageFromWs.ToLoginName)
			newMessageFromWs.FromLoginName = AllConnection.ActiveCLient[ws].LoginName
			newMessageFromWs.TimeStamp = time.Now()
			Connect.newMassege <- &newMessageFromWs
			var id string
			id = NexusIsExixst(newMessageFromWs.FromLoginName, newMessageFromWs.ToLoginName)
			if id == "NOT FOUND" {
				id = UuidGenerator()
				AddMessageClient(id, &newMessageFromWs, dbconnect.DatabaseConnectionMessage())
				AddConversations(newMessageFromWs.FromLoginName, newMessageFromWs.ToLoginName, id)

			} else {
				AddMessageClient(id, &newMessageFromWs, dbconnect.DatabaseConnectionMessage())
			}

		}

	}
}
func sendMessage(ws *websocket.Conn, message *model.Message) {
	websocket.JSON.Send(ws, message)
}
