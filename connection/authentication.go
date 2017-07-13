/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package connection

import (
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/ernestio/ernest-sdk/models"
)

// Authenticate : authenticate against an ernest instance
func (c *Conn) Authenticate() error {
	var auth models.Authentication

	form := url.Values{}
	form.Add("username", c.config.Username)
	form.Add("password", c.config.Password)

	reader := strings.NewReader(form.Encode())

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	resp, err := c.Post("/auth", "application/x-www-form-urlencoded", data)
	if err != nil {
		return err
	}

	err = ReadJSON(resp.Body, &auth)
	if err != nil {
		return err
	}

	c.config.Token = auth.Token

	return nil
}
