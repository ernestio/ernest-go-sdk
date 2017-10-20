/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

// Project : stores project data
type Project struct {
	ID           int                    `json:"id"`
	Name         string                 `json:"name"`
	Type         string                 `json:"type"`
	Credentials  map[string]interface{} `json:"credentials,omitempty"`
	Environments []string               `json:"environments,omitempty"`
	Roles        []string               `json:"roles,omitempty"`
}
