/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package builds

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

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
	default:
		w.WriteHeader(404)
	}
}

func handleget(w http.ResponseWriter, r *http.Request) {
	s := `[{"id":"1", "status":"completed"},{"id":"2", "status":"running"}]`

	if rpath(r.URL) == "/api/projects/test/envs/test/builds/1" {
		s = `{"id":"1", "status":"completed"}`
	}

	w.Write([]byte(s))
}

func handlepost(w http.ResponseWriter, r *http.Request) {
	m := models.Build{
		ID:        "1",
		Type:      "apply",
		Status:    "running",
		CreatedAt: "2017-01-01",
	}

	data, err := json.Marshal(m)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(201)
	w.Write(data)
}
