/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package reports

import "github.com/ernestio/ernest-go-sdk/connection"

var apiroute = "/api/reports/"

// Reports ...
type Reports struct {
	Conn *connection.Conn
}
