/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package generic

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ernestio/ernest-go-sdk/connection"
	"github.com/ernestio/ernest-go-sdk/models"
)

// Update : updates a generic model
func (u *Generic) Update(m models.Generic) error {
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}

	path := fmt.Sprintf("%s%s", u.APIRoute, m.GetID())

	resp, err := u.Conn.Put(path, "application/json", data)
	if err != nil {
		return err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println(err.Error())
		}
	}()

	return connection.ReadJSON(resp.Body, m)
}
