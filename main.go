package main

import (
	"flag"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	// "github.com/moazrefat/bankapp/pkg/login"
	login "./pkg/login"
)

func main() {
	fmt.Println("bank-app is starting ... ")
	var portNum = flag.String("p", "80", "Specify application server listening port")
	flag.Parse()
	fmt.Println("Vulnapp server listening : " + *portNum)

	http.HandleFunc("/", welcome)
	http.HandleFunc("/login", login.Login)
	// http.HandleFunc("/welcomepage", welcomepage)
}

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("this is welcome message")
}

// func welcomePage() (w http.ResponseWriter, r *http.Request) {
// 	userName, sessionID, userID, err := cookie.GetUserIDFromCookie(r)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }
