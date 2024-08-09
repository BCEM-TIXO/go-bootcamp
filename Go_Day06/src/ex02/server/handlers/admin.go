package handlers

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"ex02/server/model"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func (a *App) Authorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := readCookie("token", r)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		if _, ok := a.auth[token]; !ok {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}
func readCookie(name string, r *http.Request) (value string, err error) {
	if name == "" {
		return value, errors.New("you are trying to read empty cookie")
	}
	cookie, err := r.Cookie(name)

	if err != nil {
		return value, err
	}
	str := cookie.Value
	value, _ = url.QueryUnescape(str)
	return value, err
}

func (a App) checkAdmin(login, password string) bool {
	return a.adminLogin.Login == login && a.adminLogin.Password == password
}

func (a App) AdminLogin(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("username")
	password := r.FormValue("password")

	if !a.checkAdmin(login, password) {
		http.Error(w, "Incorrect login:pass", http.StatusInternalServerError)
		return
	}
	//логин и пароль совпадают, поэтому генерируем токен, пишем его в кеш и в куки
	time64 := time.Now().Unix()
	timeInt := fmt.Sprint(time64)
	token := login + password + timeInt
	hashToken := md5.Sum([]byte(token))
	hashedToken := hex.EncodeToString(hashToken[:])
	a.auth[hashedToken] = true
	livingTime := 60 * time.Minute
	expiration := time.Now().Add(livingTime)
	//кука будет жить 1 час
	cookie := http.Cookie{Name: "token", Value: url.QueryEscape(hashedToken), Expires: expiration}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func (a App) AdminPage(w http.ResponseWriter, r *http.Request) {
	if err := a.Templates["adminPost"].Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (a App) AdminLoginPage(w http.ResponseWriter, r *http.Request) {
	if err := a.Templates["adminLogin"].Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (a *App) AdminPost(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	file := r.FormValue("file")
	file += ".md"
	file = a.articlePath + file
	err := a.Converter.MdToHTML(file, title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	record := model.ArticleRecord{Title: title}
	err = a.DB.CreateArticleRecord(record)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	a.AdminPage(w, r)
}
