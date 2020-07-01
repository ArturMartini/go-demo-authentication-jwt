package middleware

import (
	"authentication/src/authentication/canonical"
	"authentication/src/authentication/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

var (
	svc service.Service
)

func Start() {
	svc = service.New()
	router := mux.NewRouter()
	router.HandleFunc("/login", login).Methods(http.MethodPost)
	http.ListenAndServe(":8080", router)
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

func logout(w http.ResponseWriter, r *http.Request) {

}
