package middleware

import (
	"encoding/json"
	"github.com/ArturMartini/go-demo-login-jwt/canonical"
	"github.com/ArturMartini/go-demo-login-jwt/jwt"
	"github.com/ArturMartini/go-demo-login-jwt/service"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	svc service.Service
)

func Start() error {
	svc = service.New()
	router := mux.NewRouter()
	router.HandleFunc("/login", login).Methods(http.MethodPost)
	router.HandleFunc("/client", demo).Methods(http.MethodPost)
	return http.ListenAndServe(":8080", router)
}

func demo(w http.ResponseWriter, r *http.Request) {
	if checkAccess(r) {
		err := svc.Demo()
		if err != nil {
			checkError(err)
		}
		w.WriteHeader(http.StatusOK)
	}

	w.WriteHeader(http.StatusForbidden)
}

func checkAccess(r *http.Request) bool {
	authorization := r.Header.Get("Authorization")
	jwtStr := strings.TrimSpace(strings.ReplaceAll(authorization, "Bearer", ""))
	token, err := jwt.Decode(jwtStr)
	if err != nil || !token.Valid {
		return false
	}
	return true
}

func login(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	login := canonical.Login{}
	err = json.Unmarshal(bytes, &login)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	jwt, err := svc.Login(login)
	if err != nil {
		checkError(err)
		return
	}

	resp, err := json.Marshal(&jwt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func checkError(err error) {
}

