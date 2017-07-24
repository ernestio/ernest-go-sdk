/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package groups

import "github.com/ernestio/ernest-go-sdk/connection"

var apiroute = "/api/groups/"

// Groups ...
type Groups struct {
	Conn *connection.Conn
}
