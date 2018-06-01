/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package builds

// Stream : returns a stream channel for all build messages
func (b *Builds) Stream(id string) (chan []byte, error) {
	return b.Conn.WSStream("/events", id)
}
