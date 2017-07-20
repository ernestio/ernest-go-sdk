/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package groups

import (
	"encoding/json"

	"github.com/ernestio/ernest-sdk/connection"
	"github.com/ernestio/ernest-sdk/models"
)

// Create : creates a group
func (g *Groups) Create(m *models.Group) error {
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}

	resp, err := g.Conn.Post(apiroute, "application/json", data)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return connection.ReadJSON(resp.Body, m)
}
