/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package loggers

import (
	"github.com/ernestio/ernest-go-sdk/client/generic"
)

// Delete : delete a logger
func (u *Loggers) Delete(t string) error {
	return generic.New(u.Conn, apiroute).Delete(t)
}
