/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package services

import (
	"github.com/ernestio/ernest-go-sdk/connection"
	"github.com/ernestio/ernest-go-sdk/models"
)

// List : list all availabile services
func (s *Services) List() ([]*models.Service, error) {
	var ms []*models.Service

	resp, err := s.Conn.Get(apiroute)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return ms, connection.ReadJSON(resp.Body, &ms)
}
