/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package loggers

import (
	"log"

	"github.com/ernestio/ernest-go-sdk/connection"
	"github.com/ernestio/ernest-go-sdk/models"
)

// List : list all availabile loggers
func (l *Loggers) List() ([]*models.Logger, error) {
	var ms []*models.Logger

	resp, err := l.Conn.Get(apiroute)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println(err.Error())
		}
	}()

	return ms, connection.ReadJSON(resp.Body, &ms)
}
