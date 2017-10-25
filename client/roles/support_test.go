/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package roles

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/ernestio/ernest-go-sdk/connection"
	"github.com/ernestio/ernest-go-sdk/models"
)

func rpath(u *url.URL) string {
	s := u.Path

	strings.Trim(s, u.Scheme)
	strings.Trim(s, u.Host)

	return s
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

func handleget(w http.ResponseWriter, r *http.Request) {
	s := `[{"id":"test-1", "user_id":"usr-1","role":"rol-1","resource_type":"res-1"},{"id":"test-2", "user_id":"usr-2","role":"rol-2","resource_type":"res-2"}]`

	if rpath(r.URL) == "/api/roles/test-1" {
		s = `{"id":"test-1", "user_id":"usr-1","role":"rol-1","resource_type":"res-1"}`
	}

	_, _ = w.Write([]byte(s))
}

func handlepost(w http.ResponseWriter, r *http.Request) {
	var m models.Role

	err := connection.ReadJSON(r.Body, &m)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	m.ID = "test-1"

	data, err := json.Marshal(m)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(201)
	_, _ = w.Write(data)
}

func handleput(w http.ResponseWriter, r *http.Request) {
	var m models.Role

	if rpath(r.URL) != "/api/roles/test-1" {
		w.WriteHeader(404)
	}

	err := connection.ReadJSON(r.Body, &m)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	data, err := json.Marshal(m)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	_, _ = w.Write(data)
}

func handledelete(w http.ResponseWriter, r *http.Request) {
	if rpath(r.URL) != "/api/roles/test-1" {
		w.WriteHeader(404)
	}
}
