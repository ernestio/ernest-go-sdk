/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package builds

import (
	"fmt"

	"github.com/ernestio/ernest-go-sdk/connection"
	"github.com/ernestio/ernest-go-sdk/models"
)

// List : list all availabile builds for a environment
func (b *Builds) List(project, environment string) ([]*models.Build, error) {
	var ms []*models.Build

	path := fmt.Sprintf(apiroute, project, environment)
	resp, err := b.Conn.Get(path)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return ms, connection.ReadJSON(resp.Body, &ms)
}
