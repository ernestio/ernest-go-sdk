/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package builds

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/ernestio/ernest-go-sdk/models"
	"github.com/gorilla/websocket"
	"github.com/r3labs/broadcast"
)

var upgrader = websocket.Upgrader{}

// Session : stores authentication data
type Session struct {
	Token         string  `json:"token"`
	Stream        *string `json:"stream"`
	EventID       *string `json:"event_id"`
	Username      string
	Authenticated bool
}

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

func unauthorized(w http.ResponseWriter) error {
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
	return errors.New("Unauthorized")
}

func authenticate(w http.ResponseWriter, c *websocket.Conn) (*Session, error) {
	var s Session

	mt, message, err := c.ReadMessage()
	if err != nil {
		return nil, badrequest(w)
	}

	err = json.Unmarshal(message, &s)
	if err != nil {
		return nil, badrequest(w)
	}

	if err != nil || s.Token == "BADTOKEN" {
		c.WriteMessage(mt, []byte(`{"status": "unauthorized"}`))
		return nil, unauthorized(w)
	}

	err = c.WriteMessage(mt, []byte(`{"status": "ok"}`))
	if err != nil {
		return nil, internalerror(w)
	}

	return &s, nil
}

func streamhandler(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		upgradefail(w)
		return
	}
	defer c.Close()

	var authorized bool
	var ch chan *broadcast.Event
	var sub *broadcast.Subscriber

	defer func() {
		if ch != nil && sub != nil {
			sub.Disconnect(ch)
		}
	}()

	for {
		if !authorized {
			areq, err := authenticate(w, c)
			if err != nil {
				return
			}

			sub, ch, err = register(w, areq)
			if err != nil {
				return
			}
		} else {
			msg := <-ch
			err := c.WriteMessage(websocket.TextMessage, msg.Data)
			if err != nil {
				internalerror(w)
				return
			}
		}
	}
}

func register(w http.ResponseWriter, s *Session) (*broadcast.Subscriber, chan *broadcast.Event, error) {
	if s.Stream == nil {
		return nil, nil, badstream(w)
	}

	if !bc.StreamExists(*s.Stream) && !bc.AutoStream {
		return nil, nil, badstream(w)
	} else if !bc.StreamExists(*s.Stream) && bc.AutoStream {
		bc.CreateStream(*s.Stream)
	}

	sub := bc.GetSubscriber(s.Username)
	if sub == nil {
		sub = broadcast.NewSubscriber(s.Username)
		bc.Register(*s.Stream, sub)
	}

	return sub, sub.Connect(), nil
}

func upgradefail(w http.ResponseWriter) {
	http.Error(w, "Unable to upgrade to websocket connection", http.StatusBadRequest)
}

func badrequest(w http.ResponseWriter) error {
	http.Error(w, "Could not process sent data", http.StatusBadRequest)
	return errors.New("Could not process sent data")
}

func badstream(w http.ResponseWriter) error {
	http.Error(w, "Please specify a valid stream", http.StatusBadRequest)
	return errors.New("Please specify a valid stream")
}

func internalerror(w http.ResponseWriter) error {
	http.Error(w, "Internal server error", http.StatusInternalServerError)
	return errors.New("Internal server error")
}
