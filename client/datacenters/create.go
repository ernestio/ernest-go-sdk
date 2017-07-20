/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package datacenters

import (
	"encoding/json"

	"github.com/ernestio/ernest-sdk/connection"
	"github.com/ernestio/ernest-sdk/models"
)

// Create : creates a datacenter
func (d *Datacenters) Create(m *models.Datacenter) error {
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}

	resp, err := d.Conn.Post(apiroute, "application/json", data)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return connection.ReadJSON(resp.Body, m)
}
