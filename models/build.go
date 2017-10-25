/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import "time"

// Build : stores build data
type Build struct {
	ID            string                 `json:"id"`
	EnvironmentID int                    `json:"environment_id"`
	UserID        int                    `json:"user_id"`
	Username      string                 `json:"user_name"`
	Type          string                 `json:"type"`
	Status        string                 `json:"status"`
	Definition    string                 `json:"definition"`
	Mapping       map[string]interface{} `json:"mapping"`
	CreatedAt     time.Time              `json:"created_at"`
	UpdatedAt     time.Time              `json:"updated_at"`
}

// GetID : get the id for the current object
func (a *Build) GetID() string {
	return a.ID
}
