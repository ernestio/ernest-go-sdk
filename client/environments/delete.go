/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package environments

import (
	"fmt"

	"github.com/ernestio/ernest-go-sdk/connection"
	"github.com/ernestio/ernest-go-sdk/models"
)

// Delete : delete a environment
func (e *Environments) Delete(project, name string) (*models.Build, error) {
	var m models.Build

	path := fmt.Sprintf(apiroute+"%s", project, name)
	resp, err := e.Conn.Delete(path)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return &m, connection.ReadJSON(resp.Body, &m)
}

// ForceDeletion : delete a environment
func (e *Environments) ForceDeletion(project, name string) (*models.Build, error) {
	var m models.Build

	path := fmt.Sprintf(apiroute+"%s/actions/force/", project, name)
	resp, err := e.Conn.Delete(path)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return &m, connection.ReadJSON(resp.Body, &m)
}
