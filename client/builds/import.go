/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package builds

import (
	"encoding/json"
	"fmt"

	"github.com/ernestio/ernest-go-sdk/connection"
	"github.com/ernestio/ernest-go-sdk/models"
)

// Import : creates a an import build for a service
func (b *Builds) Import(service string, m *models.Import) (*models.Build, error) {
	var bm models.Build

	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("/api/services/%s/import/", service)

	resp, err := b.Conn.Post(path, "application/json", data)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return &bm, connection.ReadJSON(resp.Body, &bm)
}
