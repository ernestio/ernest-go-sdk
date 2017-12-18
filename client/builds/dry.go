/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package builds

import (
	"fmt"

	"github.com/ernestio/ernest-go-sdk/connection"
	"github.com/ernestio/ernest-go-sdk/models"
)

// Dry : simulates a build creation
func (b *Builds) Dry(definition []byte) (*[]string, error) {
	var m []string
	var d models.Definition

	if err := d.Load(definition); err != nil {
		return nil, err
	}

	path := fmt.Sprintf(apiroute, d.Project, d.Name)

	resp, err := b.Conn.Post(path+"/?dry=true", "application/yaml", definition)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return &m, connection.ReadJSON(resp.Body, &m)
}
