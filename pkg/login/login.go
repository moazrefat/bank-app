package login

import (
	"database/sql"
	"encoding/base64"
	"log"
	"net/http"
	"text/template"

	"golang.org/x/crypto/bcrypt"

	// cookie "../cookie"
	"github.com/moazrefat/bankapp/pkg/cookie"
)

type Person struct {
	UserName string
	Name     string
	Email    string
	Age      int
}

func isZeroString(st string) bool {
	if len(st) == 0 {
		return false
	}
	return true
}

func SearchID(email string) int {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/bankapp?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	sql := "select id from bankapp.user where email=?"
	res, err := db.Query(sql, email)
	if err != nil {
		log.Println(err)
	}
	var id int
	for res.Next() {
		err = res.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("ID,", id)
	}
	return id
}

func LoginedHome(w http.ResponseWriter, r *http.Request) {

	_, _, userID, err := cookie.GetUserIDFromCookie(r)
	if err != nil {
		log.Println(err)
	}
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/bankapp?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	sql := "select name,email,age from bankapp.user where id=?"
	res, err := db.Query(sql, userID)
	if err != nil {
		log.Println(err)
	}
	var Age int
	var Name, Email string
	for res.Next() {
		err := res.Scan(&Name, &Email, &Age)
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Printf("name:%v, Email:%v, Age:%v", Name, Email, Age)
	var data Person
	data.Name = Name
	data.Age = Age
	data.Email = Email
	t, _ := template.ParseFiles("./views/public/loginedhome.gtpl")
	t.Execute(w, data)
}

func LogoutedHome(w http.ResponseWriter, r *http.Request) {
	_, sessionID, userID, err := cookie.GetUserIDFromCookie(r)
	if err != nil {
		log.Println(err)
	}
	if cookie.CheckSessionsCount(userID, sessionID) {
		StoreSID(userID, sessionID)
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

func BcryptPasswd(passwd string) []byte {
	bcryptPasswd, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	// fmt.Printf("password %v, bycrpy %v\n", passwd, string(bcryptPasswd))
	return bcryptPasswd
}

func BcryptValidation(id int, passwd string) bool {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/bankapp?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	sql := "select passwd from bankapp.user where id=?"
	res, err := db.Query(sql, id)
	if err != nil {
		log.Println(err)
	}

	var hashFromDatabase string
	for res.Next() {
		err := res.Scan(&hashFromDatabase)
		if err != nil {
			log.Fatal(err)
		}
	}
	// log.Println("hash", hashFromDatabase)

	// bcryptPasswd := BcryptPasswd(passwd)

	err = bcrypt.CompareHashAndPassword([]byte(hashFromDatabase), []byte(passwd))
	if err != nil {
		log.Println("login failed")
		return false
	}

	log.Printf("Successful login for ID %v ", id)
	return true
}

func CheckPasswd(id int, passwd string) (string, bool) {
	// passwdStatus := BcryptValidation(passwd)
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/bankapp?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	sql := "select name from bankapp.user where id=?"
	res, err := db.Query(sql, id)
	if err != nil {
		log.Println(err)
	}
	var name string
	for res.Next() {
		err := res.Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Name:", name)
	}
	return name, BcryptValidation(id, passwd)
}

func StoreSID(uid int, sid string) {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/bankapp?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	if cookie.CheckSessionsCount(uid, sid) {
		_, err = db.Exec("insert into bankapp.sessions(uid,sessionid) value (?,?)", uid, sid)
		if err != nil {
			log.Println(err)
		}
	} else {
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	log.Println("method,", r.Method)

	if r.Method == "GET" {
		if cookie.CheckSessionID(r) {
			http.Redirect(w, r, "/home", 302)
		} else {
			t, _ := template.ParseFiles("./views/public/login.gtpl")
			t.Execute(w, nil)
		}
	} else if r.Method == "POST" {
		r.ParseForm()
		if isZeroString(r.FormValue("email")) && isZeroString(r.FormValue("passwd")) {
			log.Println("passwd", r.Form["passwd"])
			log.Println("email", r.Form["email"])

			email := r.FormValue("email")
			id := SearchID(email)
			if id != 0 {
				passwd := r.FormValue("passwd")
				name, passwordStatus := CheckPasswd(id, passwd)
				log.Println("passwordStatus", passwordStatus)
				if name != "" && passwordStatus {
					// fmt.Println(name)
					t, _ := template.ParseFiles("./views/public/logined.gtpl")
					encodeEmail := base64.StdEncoding.EncodeToString([]byte(email))
					log.Println("encodeEmail", encodeEmail)
					cookieSID := &http.Cookie{
						Name:  "SessionID",
						Value: encodeEmail,
					}
					cookieUserName := &http.Cookie{
						Name:  "UserName",
						Value: name,
					}
					StoreSID(id, encodeEmail)
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
