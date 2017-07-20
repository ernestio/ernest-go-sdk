/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

// User : stores user data
type User struct {
	ID       int    `json:"id"`
	GroupID  int    `json:"group_id"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
	Salt     string `json:"salt,omitempty"`
	Admin    bool   `json:"admin"`
}
