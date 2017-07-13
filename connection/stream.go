/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package connection

import "github.com/r3labs/sse"

// Stream : connects to an sse stream, returns a channel
func (c *Conn) Stream(id string) (chan *sse.Event, error) {
	return nil, nil
}
