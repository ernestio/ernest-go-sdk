/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package loggers

import (
	"github.com/ernestio/ernest-go-sdk/client/generic"
	"github.com/ernestio/ernest-go-sdk/models"
)

// List : list all availabile loggers
func (l *Loggers) List() (ms []*models.Logger, err error) {
	err = generic.New(l.Conn, apiroute).List(&ms)
	return
}
