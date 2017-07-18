/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package services

import (
	"fmt"

	"github.com/ernestio/ernest-sdk/connection"
	"github.com/ernestio/ernest-sdk/models"
)

// Delete : delete a service
func (s *Services) Delete(name string) (*models.Service, error) {
	var m models.Service

	path := fmt.Sprintf("%s%s", apiroute, name)
	resp, err := s.Conn.Delete(path)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return &m, connection.ReadJSON(resp.Body, &m)
}
