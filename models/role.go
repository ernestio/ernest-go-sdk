/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

// Role : stores user data
type Role struct {
	ID       string `json:"resource_id"`
	User     string `json:"user_id"`
	Role     string `json:"role"`
	Resource string `json:"resource_type"`
}

// GetID : get the id for the current object
func (n *Role) GetID() string {
	return n.ID
}
