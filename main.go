package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"./pkg/cookie"
	_ "github.com/go-sql-driver/mysql"

	// login "./pkg/login"
	// logout "./pkg/logout"
	// register "./pkg/register"
	"github.com/moazrefat/bankapp/pkg/login"
	"github.com/moazrefat/bankapp/pkg/logout"
	"github.com/moazrefat/bankapp/pkg/register"
)

type Person struct {
	UserName string
}

func main() {
	log.Println("BankApp is starting ... ")
	var portNum = flag.String("p", "80", "Specify application server listening port")
	flag.Parse()

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	http.HandleFunc("/", login.LogoutedHome)
	http.HandleFunc("/login", login.Login)
	http.HandleFunc("/logout", logout.Logout)
	http.HandleFunc("/profile", Profile)
	http.HandleFunc("/home", login.LoginedHome)
	http.HandleFunc("/register", register.NewUserRegister)
	log.Println("bankapp server listening : " + *portNum)
	err := http.ListenAndServe(":"+*portNum, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// func Welcome(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("this is welcome message")
// }

func Profile(w http.ResponseWriter, r *http.Request) {
	userName, sessionID, userID, err := cookie.GetUserIDFromCookie(r)
	if err != nil {
		log.Println(err)
	}
	if cookie.CheckSessionsCount(userID, sessionID) {
		login.StoreSID(userID, sessionID)
	} else {
		log.Println("not register sessionID")
	}

	if sessionID == "" {
		log.Println("sid not exist")
		t, _ := template.ParseFiles("./views/public/error.gtpl")
		t.Execute(w, nil)
	} else {
		if r.Method == "GET" {

			if userID != 0 {
				uid := strconv.Itoa(userID)
				cookieUserID := &http.Cookie{
					Name:  "UserID",
					Value: uid,
				}

				deleAdminID := &http.Cookie{
					Name:  "adminSID",
					Value: "",
				}

				http.SetCookie(w, cookieUserID)
				http.SetCookie(w, deleAdminID)
				p := Person{UserName: userName}
				t, _ := template.ParseFiles("./views/public/profile.gtpl")
				t.Execute(w, p)
			}

		} else {
			http.NotFound(w, r)
		}
	}
}
