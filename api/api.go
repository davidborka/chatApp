package api

import (
	"log"
	"net/http"

	"github.com/davidborka/chatApp/api/auth"
	"github.com/davidborka/chatApp/api/handlers"
	"github.com/julienschmidt/httprouter"
)

func StartServer() {
	auth.SigningKey, auth.VerifyKey = auth.InitKeys()

	router := httprouter.New()
	router.NotFound = http.FileServer(http.Dir("public/"))
	//router.ServeFiles("/public/*filepath", http.Dir("./public"))

	router.POST("/register", handlers.RegisterNewClient)
	router.POST("/login", handlers.LoginHandler)
	router.GET("/protected", handlers.Proteced)
	router.GET("/logout", handlers.Logout)
	router.GET("/chat", handlers.Websocket)
	//List Message
	/*
		router.GET("/message", GetMessageToClient)
		router.GET("/message/:loginname", GetMessageToClient)
	*/
	if err := http.ListenAndServe(":9000", router); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
