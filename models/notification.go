/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

// Notification : stores user data
type Notification struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Config  string `json:"config"`
	Members string `json:"members"`
}

// GetID : get the id for the current object
func (n *Notification) GetID() string {
	return n.Name
}
