/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

// Environment : stores environment data
type Environment struct {
	ID          int                    `json:"id"`
	ProjectID   int                    `json:"project_id"`
	Project     string                 `json:"project,omitempty"`
	Provider    string                 `json:"provider,omitempty"`
	Name        string                 `json:"name"`
	Type        string                 `json:"type"`
	Status      string                 `json:"status"`
	Options     map[string]interface{} `json:"options,omitempty"`
	Credentials map[string]interface{} `json:"credentials,omitempty"`
	Roles       []string               `json:"roles,omitempty"`
}
