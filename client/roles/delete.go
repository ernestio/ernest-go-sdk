/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package roles

import (
	"encoding/json"
	"log"

	"github.com/ernestio/ernest-go-sdk/models"
)

// Delete : delete a role
func (u *Roles) Delete(m *models.Role) error {
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}

	resp, err := u.Conn.Request("DELETE", apiroute, "application/json", data, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println(err.Error())
		}
	}()

	return nil
}
