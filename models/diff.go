/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

// Diff : diff request
type Diff struct {
	From string `json:"from_id,omitempty"`
	To   string `json:"to_id,omitempty"`
}
