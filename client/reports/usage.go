/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package reports

import (
	"io/ioutil"
	"log"
)

// Usage : get an usage report
func (l *Reports) Usage(from, to string) ([]byte, error) {
	path := apiroute + "usage/?from=" + from + "&to=" + to
	resp, err := l.Conn.Get(path)
	if err != nil {
		return []byte(""), err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println(err.Error())
		}
	}()

	return ioutil.ReadAll(resp.Body)
}
