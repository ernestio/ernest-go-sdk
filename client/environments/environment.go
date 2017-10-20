/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package environments

import "github.com/ernestio/ernest-go-sdk/connection"

var apiroute = "/api/projects/%s/envs/"

// Environments ...
type Environments struct {
	Conn *connection.Conn
}
