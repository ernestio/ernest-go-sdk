/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package environments

import (
	"fmt"

	"github.com/ernestio/ernest-go-sdk/connection"
	"github.com/ernestio/ernest-go-sdk/models"
)

// List : list all availabile environments for a project
func (e *Environments) List(project string) ([]*models.Environment, error) {
	var ms []*models.Environment

	resp, err := e.Conn.Get(fmt.Sprintf(apiroute, project))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return ms, connection.ReadJSON(resp.Body, &ms)
}

// ListAll : list all availabile environments
func (e *Environments) ListAll() ([]*models.Environment, error) {
	var ms []*models.Environment

	resp, err := e.Conn.Get("/api/envs/")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return ms, connection.ReadJSON(resp.Body, &ms)
}
