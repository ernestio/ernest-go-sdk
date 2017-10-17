/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package environments

import (
	"encoding/json"
	"fmt"

	"github.com/ernestio/ernest-go-sdk/connection"
	"github.com/ernestio/ernest-go-sdk/models"
)

// Update : updates a environment
func (e *Environments) Update(m *models.Environment) error {
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}

	path := fmt.Sprintf("%s%s", apiroute, m.Name)

	resp, err := e.Conn.Put(path, "application/json", data)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return connection.ReadJSON(resp.Body, m)
}
