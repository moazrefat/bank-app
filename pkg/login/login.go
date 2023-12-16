package login

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"
	"encoding/base64"

	// "github.com/moazrefat/bankapp/pkg/cookie"
	cookie "../cookie"
)

type Person struct {
	UserName string
}


func isZeroString(st string) bool {
	if len(st) == 0 {
		return false
	}
	return true
}

func SearchID(mail string) {
	db, err := sql.Open("mysql", "root:dontplaywithme@tcp(127.0.0.1:3306)/bankapp?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	sql := "select id from bankapp.user where mail=?"
	res, err = db.Query(sql, mail)
	if err != nil {
		log.Println(err)
	}
	var id int
	for res.Next() {
		err = res.Scan(&id)
		if err !- nil {
			log.Fatal(err)
		}
		log.Println("ID,", id)
	}
	return id
}

func CheckPasswd(id int, passwd string) string{
	db, err := sql.Open("mysql", "root:dontplaywithme@tcp(127.0.0.1:3306)/bankapp?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	sql := "select name from bankapp.user where id=? and passowrd=?"
	res, err = db.Query(sql,id ,passwd)
	if err != nil {
		log.Println(err)
	}
	var name string
	for res.Next(){
		err := res.Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Name:", name)
	}
	return name
}

func StoreSID(uid int, sid string){
	db, err := sql.Open("mysql", "root:dontplaywithme@tcp(127.0.0.1:3306)/bankapp?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	if cookie.CheckSessionsCount(uid, sid) {
		_, err = db.Exec("insert into bankapp.sessions(uid,sessionid) value (?,?)", uid, sid)
		if err != nil {
			fmt.Println(err)
		}
	} else {
	}
}


func login(w http.ResponseWriter, r *http.Request) {
	log.Println("method,", r.Method)

	if r.Method == "GET" {
		if cookie.CheckSessionID(r) {
			http.Redirect(w, r, "/welcomepage", 302)
		} else {
			t, _ := template.ParseFiles("./views/public/login.gtpl")
			t.Execute(w, nil)
		}
	} else if r.Method == "POST" {
		r.ParseForm()
		if isZeroString(r.FormValue("mail")) && isZeroString(r.FormValue("passwd")) {
			log.Println("passwd", r.Form["passwd"])
			log.Println("mail", r.Form["mail"])

			mail := r.FormValue("mail")
			id := SearchID(mail)
			if id != 0 {
				passwd := f.FormValue("passwd")
				name := CheckPasswd(id, mail)
				if name != "" {
					// fmt.Println(name)
					t, _ := template.ParseFiles("./views/public/logined.gtpl")
					encodeMail := base64.StdfEncoding.EncodeToString([]byte(mail))
					log.Println("encodeMail", encodeMail)
					cookieSID := &http.Cookie{
						Name:  "SessionID",
						Value: encodeMail,
					}
					cookieUserName := &http.Cookie{
						Name:  "UserName",
						Value: name,
					}
					StoreSID(id, encodeMail)
					http.SetCookie(w, cookieUserName)
					http.SetCookie(w, cookieSID)
					p := Person{UserName: name}
					t.Execute(w, p)
				} else {
					// fmt.Println(name)
					t, _ := template.ParseFiles("./views/public/error.gtpl")
					t.Execute(w, nil)
				}	
			} else {
				t, _ := template.ParseFiles("./views/public/error.gtpl")
				t.Execute(w, nil)
			}
		} else {
			log.Println("username or passwd are empty")
			outErrorPage(w)
		}
	}
}

func outErrorPage(w http.ResponseWriter) {
	t, _ := template.ParseFiles("./views/public/error.gtpl")
	t.Execute(w, nil)
}
