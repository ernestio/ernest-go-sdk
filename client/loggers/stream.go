/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package loggers

import "github.com/r3labs/sse"

// Stream : returns a stream channel for all build messages
func (l *Loggers) Stream(id string) (chan *sse.Event, error) {
	return l.Conn.Stream("/logs", id)
}
