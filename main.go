package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	// "github.com/moazrefat/bankapp/pkg/login"
	login "./pkg/login"
)

func main() {
	fmt.Println("bank-app is starting ... ")
	var portNum = flag.String("p", "80", "Specify application server listening port")
	flag.Parse()

	http.HandleFunc("/", Welcome)
	http.HandleFunc("/login", login.Login)
	fmt.Println("bankapp server listening : " + *portNum)
	err := http.ListenAndServe(":"+*portNum, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// http.HandleFunc("/welcomepage", welcomepage)

func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("this is welcome message")
}

// func welcomePage() (w http.ResponseWriter, r *http.Request) {
// 	userName, sessionID, userID, err := cookie.GetUserIDFromCookie(r)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }
