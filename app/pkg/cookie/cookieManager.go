package cookie

import (
	"database/sql"
	"encoding/base64"
	"errors"
	"log"
	"net/http"
	"strconv"
)

func CheckSessionsCount(uid int, sid string) bool {
	db, err := sql.Open("mysql", "user:password@tcp(database.bankapp.svc:3306)/bankapp?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("uid : ", uid)

	var count int

	if err := db.QueryRow("select count(sessionid) from bankapp.sessions where uid=?", uid).Scan(&count); err != nil {
		log.Println(err)
	}
	log.Println("connecting to database ...")
	log.Println("count is : ", count)

	if count != 0 {
		return false
	}
	return true

}

func ValidateCorrectCookie(uid int, sid string) bool {
	db, err := sql.Open("mysql", "user:password@tcp(database.bankapp.svc:3306)/bankapp?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Println("connecting to database ...")
	var sessionID string
	if err := db.QueryRow("select sessionid from bankapp.sessions where uid=?", uid).Scan(&sessionID); err != nil {
		log.Println(err)
	}

	if sessionID == sid {
		return true
	} else {
		return false
	}

}

func CheckSessionID(r *http.Request) bool {
	sessionID, err := r.Cookie("SessionID")
	if err != nil {
		log.Println(err)
		return false
	}

	userID, err := r.Cookie("UserID")
	if err != nil {
		log.Println(err)
		return false
	}

	if sessionID.Value == "" || userID.Value == "" {
		return false
	}

	uid, err := strconv.Atoi(userID.Value)
	if err != nil {
		log.Println(err)
	}

	if ValidateCorrectCookie(uid, sessionID.Value) {
		return true
	} else {
		return false
	}

}

func GetCookieValue(r *http.Request) (SessionID string, UserName string, UserID int, err error) {
	sessionID, err := r.Cookie("SessionID")
	if err != nil {
		log.Println(err)
		return "", "", 0, err
	}

	userID, err := r.Cookie("UserID")
	if err != nil {
		log.Println(err)
	}
	uid, err := strconv.Atoi(userID.Value)
	if err != nil {
		log.Println(err)
		return "", "", 0, err
	}

	userName, err := r.Cookie("UserName")
	if err != nil {
		log.Println(err)
		return "", "", 0, err
	}

	return sessionID.Value, userName.Value, uid, nil
}

func CheckCookieOnlyLogin(r *http.Request) (userNameCookie string, sessionIDCookie string, err error) {
	userName, err := r.Cookie("UserName")
	if err != nil {
		log.Println(err)
	}

	sessionID, err := r.Cookie("SessionID")
	if err != nil {
		log.Println(err)
	}

	log.Println(userName, sessionID)

	if userName.Value == "" && sessionID.Value == "" {
		return "", "", errors.New("Cookie not exsit")
	} else {
		return userName.Value, sessionID.Value, nil
	}
}

func GetUserIDFromCookie(r *http.Request) (userNameCookie string, sessionIDCookie string, userIDCookie int, err error) {
	userName, err := r.Cookie("UserName")
	if err != nil {
		log.Println(err)
	}

	sessionID, err := r.Cookie("SessionID")
	if err != nil {
		log.Println(err)
	}

	if userName.Value == "" && sessionID.Value == "" {
		return "", "", 0, errors.New("not exist cookie")
	} else {
		decodeEmail, err := base64.StdEncoding.DecodeString(sessionID.Value)
		if err != nil {
			log.Println(err)
		}
		email := string(decodeEmail)
		log.Println(email)

		db, err := sql.Open("mysql", "user:password@tcp(database.bankapp.svc:3306)/bankapp?parseTime=true")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		var userID int
		if err := db.QueryRow("select id from user where email=?", email).Scan(&userID); err != nil {
			log.Println("no set :", err)
		}

		log.Println(userID)
		return userName.Value, sessionID.Value, userID, nil
	}
}
