/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package users

import (
	"encoding/json"

	"github.com/ernestio/ernest-go-sdk/connection"
	"github.com/ernestio/ernest-go-sdk/models"
)

// Create : creates a user
func (u *Users) Create(m *models.User) error {
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}

	resp, err := u.Conn.Post(apiroute, "application/json", data)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return connection.ReadJSON(resp.Body, m)
}
