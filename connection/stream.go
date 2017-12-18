/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package connection

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"

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
