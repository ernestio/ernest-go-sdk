package connection

import (
	"io/ioutil"
	"net/http"
)

func handleget(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"status":"ok"}`))
}

func handlepost(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	w.WriteHeader(201)
	w.Write(data)
}

func handleput(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	w.WriteHeader(202)
	w.Write(data)
}

func handledelete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"status":"ok"}`))
}

func handleauth(w http.ResponseWriter, r *http.Request) {
	u := r.FormValue("username")
	p := r.FormValue("password")

	if u != "user" && p != "pass" {
		w.WriteHeader(403)
		return
	}

	w.Write([]byte(`{"token":"test-token"}`))
}

func testhandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleget(w, r)
	case "POST":
		handlepost(w, r)
	case "PUT":
		handleput(w, r)
	case "DELETE":
		handledelete(w, r)
	}
}
