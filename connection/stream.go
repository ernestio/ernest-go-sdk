/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package connection

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
	"github.com/r3labs/sse"
)

// Stream : connects to an sse stream, returns a channel
func (c *Conn) Stream(path string, stream string) (chan *sse.Event, error) {
	ch := make(chan *sse.Event)

	u, err := url.Parse(c.config.Target)
	if err != nil {
		return nil, err
	}

	u.Path = path

	srv := sse.NewClient(u.String())
	srv.EncodingBase64 = true
	srv.Connection.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	srv.Headers["Authorization"] = fmt.Sprintf("Bearer %s", c.config.Token)

	return ch, srv.SubscribeChan(stream, ch)
}

// WSStream : connects to a websocket stream, returns a channel
func (c *Conn) WSStream(path string, stream string) (chan []byte, error) {
	var authresp struct {
		Status string `json:"status"`
	}

	u := url.URL{Scheme: c.config.WSScheme(), Host: c.config.Hostname(), Path: path}

	websocket.DefaultDialer.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	ws, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, err
	}

	auth := fmt.Sprintf(`{"token": "%s", "stream": "%s"}`, c.config.Token, stream)

	err = ws.WriteMessage(websocket.TextMessage, []byte(auth))
	if err != nil {
		return nil, err
	}

	err = ws.ReadJSON(&authresp)
	if err != nil {
		return nil, err
	}

	if authresp.Status != "ok" {
		return nil, errors.New(authresp.Status)
	}

	events := make(chan []byte)

	go func(ws *websocket.Conn, events chan []byte) {
		defer ws.Close()
		for {
			_, ev, err := ws.ReadMessage()
			if err != nil {
				close(events)
				return
			}
			events <- ev
		}
	}(ws, events)

	return events, nil
}
