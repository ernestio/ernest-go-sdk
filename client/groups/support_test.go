/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package groups

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/ernestio/ernest-sdk/connection"
	"github.com/ernestio/ernest-sdk/models"
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
	case "DELETE":
		handledelete(w, r)
	}
}

func handleget(w http.ResponseWriter, r *http.Request) {
	s := `[{"id":1, "name":"test-1"},{"id":2, "name":"test-2"}]`

	if rpath(r.URL) == "/api/groups/1" {
		s = `{"id":1, "name":"test-1"}`
	}

	w.Write([]byte(s))
}

func handlepost(w http.ResponseWriter, r *http.Request) {
	var m models.Group

	err := connection.ReadJSON(r.Body, &m)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	m.ID = 1

	data, err := json.Marshal(m)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(201)
	w.Write(data)
}

func handledelete(w http.ResponseWriter, r *http.Request) {
	if rpath(r.URL) != "/api/groups/1" {
		w.WriteHeader(404)
		return
	}
}
