/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package connection

import (
	"encoding/json"

	"github.com/ernestio/ernest-go-sdk/models"
)

// Authenticate : authenticate against an ernest instance
func (c *Conn) Authenticate() error {
	var auth models.Authentication

	var authreq struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	authreq.Username = c.config.Username
	authreq.Password = c.config.Password

	data, err := json.Marshal(authreq)
	if err != nil {
		return err
	}

	resp, err := c.Post("/auth", "application/json", data)
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
