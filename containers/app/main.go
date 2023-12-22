package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	// cookie "./pkg/cookie"
	// login "./pkg/login"
	// logout "./pkg/logout"
	// register "./pkg/register"

	"github.com/moazrefat/bankapp/pkg/login"
	"github.com/moazrefat/bankapp/pkg/logout"
	"github.com/moazrefat/bankapp/pkg/register"
	"github.com/moazrefat/bankapp/pkg/user"
)

type Person struct {
	UserName string
}

func main() {
	log.Println("BankApp is starting ... ")
	var portNum = flag.String("p", "80", "Specify application server listening port")
	flag.Parse()

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	http.HandleFunc("/", user.LogoutedHome)
	// http.HandleFunc("/", Welcome)
	http.HandleFunc("/login", login.Login)
	http.HandleFunc("/logout", logout.Logout)
	http.HandleFunc("/profile", user.Profile)
	http.HandleFunc("/home", user.LoginedHome)
	http.HandleFunc("/register", register.NewUserRegister)
	log.Println("BankApp server listening : " + *portNum)
	err := http.ListenAndServe(":"+*portNum, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	log.Println("this is home page")
	t, _ := template.ParseFiles("./views/public/welcome.gtpl")
	t.Execute(w, nil)
}
