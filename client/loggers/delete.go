/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package loggers

import (
	"fmt"
	"log"
)

// Delete : delete a logger
func (u *Loggers) Delete(t string) error {
	path := fmt.Sprintf(apiroute+"%s", t)

	resp, err := u.Conn.Delete(path)
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
