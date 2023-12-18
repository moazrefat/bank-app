package register

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

type Person struct {
	UserName string
	Uid      int
}

func CheckUserDeplicate(email string) bool {
	db, err := sql.Open("mysql", "root:dontplaywithme@tcp(127.0.0.1:3306)/bankapp")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var count int
	if err = db.QueryRow("select count(email) from user where mail=?", email).Scan(&count); err != nil {
		fmt.Println("db error : ", err)
	} else {
		fmt.Println(reflect.TypeOf(count))
		fmt.Println(count)
	}

	if count == 0 {
		return true
	}

	return false
}

func RegisterUser(r *http.Request) bool {
	db, err := sql.Open("mysql", "root:dontplaywithme@tcp(127.0.0.1:3306)/bankapp")
	if err != nil {
		log.Fatal(err)
	}

	age, err := strconv.Atoi(r.FormValue("age"))
	if err != nil {
		fmt.Println(err)
		return false
	}
	hashPasswd, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("passwd")), bcrypt.MinCost)
	_, err = db.Exec("insert into user (name,email,age,passwd) value(?,?,?,?)", r.FormValue("name"), r.FormValue("email"), age, string(hashPasswd))
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true

}

func NewUserRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./views/public/register.gtpl")
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		r.ParseForm()
		log.Println(r.FormValue("email"))
		log.Println(r.FormValue("name"))
		log.Println(r.FormValue("age"))
		log.Println(r.FormValue("passwd"))

		if CheckUserDeplicate(r.FormValue("email")) {

			if RegisterUser(r) {

				name := r.FormValue("name")
				email := r.FormValue("email")
				encodeEmail := base64.StdEncoding.EncodeToString([]byte(email))
				fmt.Println("register successful!!")
				cookieSID := &http.Cookie{
					Name:  "SessionID",
					Value: encodeEmail,
				}
				cookieUserName := &http.Cookie{
					Name:  "UserName",
					Value: name,
				}

				http.SetCookie(w, cookieSID)
				http.SetCookie(w, cookieUserName)
				p := Person{UserName: name}
				t, _ := template.ParseFiles("./views/public/success_register.gtpl")
				t.Execute(w, p)
			}
		} else {
			t, _ := template.ParseFiles("./views/public/register_error.gtpl")
			t.Execute(w, nil)
		}

	} else {
		http.NotFound(w, r)
	}

}
