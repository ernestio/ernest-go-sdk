/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package builds

import (
	"fmt"

	"github.com/ernestio/ernest-go-sdk/connection"
	"github.com/ernestio/ernest-go-sdk/models"
)

// Get : get a service's build
func (b *Builds) Get(service, id string) (*models.Build, error) {
	var m models.Build

	path := fmt.Sprintf(apiroute+"%s", service, id)
	resp, err := b.Conn.Get(path)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return &m, connection.ReadJSON(resp.Body, &m)
}
