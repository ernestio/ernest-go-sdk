/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package generic

import (
	"fmt"
	"log"

	"github.com/ernestio/ernest-go-sdk/connection"
)

// Get : get a notification
func (l *Generic) Get(id string, m interface{}) (err error) {
	path := fmt.Sprintf(l.APIRoute + id)
	resp, err := l.Conn.Get(path)
	if err != nil {
		return err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println(err.Error())
		}
	}()

	return connection.ReadJSON(resp.Body, &m)
}
