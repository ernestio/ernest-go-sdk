/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package builds

import (
	"fmt"

	"github.com/ernestio/ernest-go-sdk/connection"
	"github.com/ernestio/ernest-go-sdk/models"
)

// Create : creates an environment build
func (b *Builds) Create(definition []byte) (*models.Build, error) {
	var m models.Build
	var d models.Definition
	var e models.Error

	err := d.Load(definition)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(apiroute, d.Project, d.Name)

	resp, err := b.Conn.Post(path, "application/yaml", definition)
	if err != nil {
		nerr := connection.ReadJSON(resp.Body, &e)
		if nerr != nil {
			return nil, err
		}

		return nil, &e
	}

	defer resp.Body.Close()

	return &m, connection.ReadJSON(resp.Body, &m)
}
