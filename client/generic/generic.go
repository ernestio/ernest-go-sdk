/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package generic

import "github.com/ernestio/ernest-go-sdk/connection"

// Generic ...
type Generic struct {
	Conn     *connection.Conn
	APIRoute string
}

// New : ..
func New(conn *connection.Conn, apiroute string) *Generic {
	g := Generic{
		Conn:     conn,
		APIRoute: apiroute,
	}
	return &g
}
