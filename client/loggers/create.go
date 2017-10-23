/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package loggers

import (
	"encoding/json"
	"log"

	"github.com/ernestio/ernest-go-sdk/connection"
	"github.com/ernestio/ernest-go-sdk/models"
)

// Create : creates a logger
func (u *Loggers) Create(m *models.Logger) error {
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}

	resp, err := u.Conn.Post(apiroute, "application/json", data)
	if err != nil {
		return err
	}

	defer func() {
		if err = resp.Body.Close(); err != nil {
			log.Println(err.Error())
		}
	}()

	return connection.ReadJSON(resp.Body, m)
}
