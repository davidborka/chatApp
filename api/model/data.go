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
	Inner struct {
		LoginName string `json:"loginname,omitempty"`
		Email     string `json:"email,omitempty"`
		Password  []byte `json:"password,omitempty"`
		Messages  []Message
	} `json:"chatapp"`
}

//Message data structure that save a database
type Message struct {
	Inner struct {
		FromLoginName string `json:"fromloginname,omitempty"`
		ToLoginName   string `json:"tologinname,omitempty"`
		TimeStamp     time.Time
		Content       string `json:"content,omitempty"`
	} `json:"newmessage"`
}

//ConnectionChan

//OnlineClient
type OnlineClient struct {
	Inner struct {
		LoginName []string `json:"name"`
	} `json:"online"`
}
