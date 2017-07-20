/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package builds

import "github.com/r3labs/sse"

// Stream : returns a stream channel for all build messages
func (b *Builds) Stream(id string) (chan *sse.Event, error) {
	return b.Conn.Stream("/events", id)
}
