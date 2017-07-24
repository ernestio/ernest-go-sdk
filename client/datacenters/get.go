/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package datacenters

import (
	"fmt"

	"github.com/ernestio/ernest-go-sdk/connection"
	"github.com/ernestio/ernest-go-sdk/models"
)

// Get : get a datacenter
func (d *Datacenters) Get(id int) (*models.Datacenter, error) {
	var m models.Datacenter

	path := fmt.Sprintf(apiroute+"%d", id)
	resp, err := d.Conn.Get(path)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return &m, connection.ReadJSON(resp.Body, &m)
}
