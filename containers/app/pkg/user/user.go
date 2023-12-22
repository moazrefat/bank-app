package user

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/moazrefat/bankapp/pkg/cookie"
	"github.com/moazrefat/bankapp/pkg/login"
)

type Person struct {
	UserName string
	Name     string
	Email    string
	Age      int
	Address  string
}

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

func LoginedHome(w http.ResponseWriter, r *http.Request) {

	_, _, userID, err := cookie.GetUserIDFromCookie(r)
	if err != nil {
		log.Println(err)
	}
	db, err := sql.Open("mysql", "user:password@tcp(database.bankapp.svc:3306)/bankapp?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	log.Println("connecting to database ...")
	var Age int
	var Name, Email, Address string
	sql := "select name,email,age from bankapp.user where id=?"
	res, err := db.Query(sql, userID)
	if err != nil {
		log.Println(err)
	}
	for res.Next() {
		err := res.Scan(&Name, &Email, &Age)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = db.QueryRow("select address from userdetails where uid=?", userID).Scan(&Address)
	if err != nil {
		log.Println(err)
	}

	log.Printf("name:%v, Email:%v, Age:%v, Address:%v", Name, Email, Age, Address)
	var data Person
	data.Name = Name
	data.Age = Age
	data.Email = Email
	data.Address = Address
	t, _ := template.ParseFiles("./views/public/loginedhome.gtpl")
	t.Execute(w, data)
}

func LogoutedHome(w http.ResponseWriter, r *http.Request) {
	_, sessionID, userID, err := cookie.GetUserIDFromCookie(r)
	if err != nil {
		log.Println(err)
	}
	if cookie.CheckSessionsCount(userID, sessionID) {
		login.StoreSID(userID, sessionID)
	} else {
		log.Println("not register sessionID")
	}

	if sessionID == "" {
		t, _ := template.ParseFiles("./views/public/logoutedhome.gtpl")
		t.Execute(w, nil)
	} else {
		if r.Method == "GET" {
			http.Redirect(w, r, "/home", 302)
		} else {
			http.NotFound(w, r)
		}
	}
}
