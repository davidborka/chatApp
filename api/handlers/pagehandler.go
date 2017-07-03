package handlers

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/websocket"

	"github.com/davidborka/chatApp/api/auth"
	"github.com/davidborka/chatApp/api/dbconnect"
	"github.com/davidborka/chatApp/api/middleware"
	"github.com/davidborka/chatApp/api/model"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

var LoginClient model.Client
var (
	LoginClients = make(map[string]model.Client)
)

func Proteced(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.Handler(validateHttp(protectedProfile)).ServeHTTP(w, r)
}
func Logout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.Handler(validateHttp(logout)).ServeHTTP(w, r)
}
func Websocket(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	websocket.Handler(HandleChatRoom).ServeHTTP(w, r)
}
func LoginHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	bucket := dbconnect.DatabaseConnection()
	fmt.Println("Login user #1")
	if _, err := bucket.Get(r.FormValue("username"), &LoginClient); err != nil {

		fmt.Println("Cant find the user")
		http.Redirect(w, r, "/", 307)
		return
	}
	if !(CompareUserPassword(LoginClient.Inner.Password, r.FormValue("password"))) {
		fmt.Println("wrong pass")
		http.Redirect(w, r, "/", 307)
		w.WriteHeader(http.StatusForbidden)
		return
	}
	expireCookie := time.Now().Add(time.Minute * 25)
	signedToken := middleware.GenerateAuthToken(LoginClient.Inner.LoginName)
	cookie := http.Cookie{Name: "Auth", Value: signedToken, Expires: expireCookie, HttpOnly: true}
	LoginClients[signedToken] = LoginClient
	http.SetCookie(w, &cookie)
	fmt.Println("setting cookie")
	http.Redirect(w, r, "/protected", 302)
}

//Middleware
func validateHttp(page http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("Auth")
		if err != nil {
			fmt.Println("not set Cookie")
			http.NotFound(w, r)
			return
		}
		fmt.Println("Login user #3")
		token, err := jwt.ParseWithClaims(cookie.Value, &middleware.Claims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method")
			}
			return auth.VerifyKey, nil
		})
		if err != nil {
			http.NotFound(w, r)
			fmt.Println("the token is not valid")
			return
		}

		if _, ok := token.Claims.(*middleware.Claims); ok && token.Valid {
			fmt.Println("Login user #4")
			page(w, r)
		} else {
			fmt.Println("something wrong with claims")
			http.NotFound(w, r)
			return
		}
	})
}
func RegisterNewClient(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	bucket := dbconnect.DatabaseConnection()
	var newClient model.Client
	clientIsOk := true
	if r.Method == "POST" {
		newClient.Inner.LoginName = r.FormValue("username")
		newClient.Inner.Email = r.FormValue("email")
	}
	password := CreatePasswordHash(r.FormValue("password"))
	newClient.Inner.Password = password
	allClient := ListAllClientFromDb(bucket)
	for _, v := range allClient {
		if newClient.Inner.LoginName == v.Inner.LoginName {
			clientIsOk = false
		}
	}
	if clientIsOk {
		if err := AddClient(newClient, bucket); !err {
			fmt.Println("Cant create the user")
		}
		//http.ServeFile(w, r, "static_pages/index.html")
	} else {
		w.Write([]byte("ERROR"))

	}

	http.Redirect(w, r, "/", 307)

}
func protectedProfile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login user #5")
	http.ServeFile(w, r, "public/profile.html")

}

// deletes the cookie
func logout(w http.ResponseWriter, r *http.Request) {
	cookies, _ := r.Cookie("Auth")
	deleteClient := LoginClients[cookies.Value]
	Connect.removeConnection <- &deleteClient
	deleteCookie := http.Cookie{Name: "Auth", Value: "none", Expires: time.Now()}
	http.SetCookie(w, &deleteCookie)

	http.Redirect(w, r, "/", 302)
	return
}
func CreatePasswordHash(password string) []byte {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// TODO: Properly handle error

		return []byte("Error")
	}
	return hash
}
func CompareUserPassword(hashedPassword []byte, password string) bool {
	if err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password)); err != nil {
		// TODO: Properly handle error

		return false
	}
	return true
}
