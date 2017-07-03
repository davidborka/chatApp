package handlers

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func StartServer() {
	go Connect.StartConnection()
	router := httprouter.New()
	router.NotFound = http.FileServer(http.Dir("public/"))
	//router.ServeFiles("/public/*filepath", http.Dir("./public"))

	router.POST("/register", RegisterNewClient)
	router.POST("/login", LoginHandler)
	router.GET("/protected", proteced)
	router.GET("/logout", logout)
	router.GET("/chat", Websocket)
	//List Message
	router.GET("/message", GetMessageToClient)
	router.GET("/message/:loginname", GetMessageToClient)
	if err := http.ListenAndServe(":9000", router); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
