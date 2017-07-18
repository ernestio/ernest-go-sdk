/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package builds

import (
	"fmt"

	"github.com/ernestio/ernest-sdk/connection"
	"github.com/ernestio/ernest-sdk/models"
)

// Create : creates a service
func (b *Builds) Create(service string, definition []byte) (*models.Build, error) {
	var m models.Build

	path := fmt.Sprintf(apiroute, service)
	resp, err := b.Conn.Post(path, "application/yaml", definition)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return &m, connection.ReadJSON(resp.Body, &m)
}
