/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package connection

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Request : make a raw request to ernest
func (c *Conn) Request(method, path, ctype string, data []byte, headers map[string]string) (*http.Response, error) {
	req, err := c.setupRequest(method, path, ctype, data, headers)
	if err != nil {
		return nil, newError(err.Error())
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Do(req)
	if err != nil {
		return resp, newError(err.Error())
	}

	return resp, responseError(resp)
}

func (c *Conn) setupRequest(method, path, ctype string, data []byte, headers map[string]string) (*http.Request, error) {
	u, err := url.Parse(c.config.Target)
	if err != nil {
		return nil, err
	}

	u.Path = path
	if strings.Contains(path, "?") {
		parts := strings.Split(path, "?")
		u.Path = parts[0]
	}

	req, err := http.NewRequest(method, u.String(), bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	if c.config.Token != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.config.Token))
	}
	req.Header.Set("Content-Type", ctype)

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	q := req.URL.Query()
	if u2, err := url.Parse("http://something.com" + path); err == nil {
		params, _ := url.ParseQuery(u2.RawQuery)
		for k, v := range params {
			q.Add(k, v[0])
		}
	}
	req.URL.RawQuery = q.Encode()

	return req, nil
}
