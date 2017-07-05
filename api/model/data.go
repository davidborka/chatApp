package model

import (
	"time"

	"golang.org/x/net/websocket"
)

//Connection
type Connection struct {
	ActiveCLient map[*websocket.Conn]Client
}

//Client data structure that save a database
type Client struct {
	Uuid      string `json:"uuid"`
	LoginName string `json:"loginname,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  []byte `json:"password,omitempty"`
}
type TypedClient struct {
	Chatapp struct {
		Uuid      string `json:"uuid"`
		LoginName string `json:"loginname,omitempty"`
		Email     string `json:"email,omitempty"`
		Password  []byte `json:"password,omitempty"`
	} `json:"chatappUsers"`
}

//Message data structure that save a database
type Message struct {
	FromLoginName string `json:"fromloginname,omitempty"`
	ToLoginName   string `json:"tologinname,omitempty"`
	TimeStamp     time.Time
	Content       string `json:"content,omitempty"`
}
type Nexus struct {
	Messages []Message
}
type Conversations struct {
	Conv []Conversation
}
type Conversation struct {
	Partner1Uuid string `json:"partneroneuuid"`
	Partner2Uuid string `json:"partnertwouuid"`
	NexusUuid    string `json:"nexusuuid"`
}

//ConnectionChan

//OnlineClient
type OnlineClient struct {
	Inner struct {
		LoginName []string `json:"name"`
	} `json:"online"`
}

type TypedMessage struct {
	Inner struct {
		FromLoginName string `json:"fromloginname,omitempty"`
		ToLoginName   string `json:"tologinname,omitempty"`
		TimeStamp     time.Time
		Content       string `json:"content,omitempty"`
	} `json:"chatapp"`
}
