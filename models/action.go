/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import "fmt"

// Action : stores environment action data
type Action struct {
	ID           int           `json:"id"`
	Type         string        `json:"type,omitempty"`
	Status       string        `json:"status,omitempty"`
	ResourceID   string        `json:"resource_id,omitempty"`
	ResourceType string        `json:"resource_type,omitempty"`
	Options      ActionOptions `json:"options,omitempty"`
}

// GetID : get the id for the current object
func (a *Action) GetID() string {
	return fmt.Sprintf("%d", a.ID)
}

// ActionOptions : stores options associated with an action
type ActionOptions struct {
	Filters     []string `json:"filters,omitempty"`
	BuildID     string   `json:"build_id,omitempty"`
	Environment string   `json:"environment,omitempty"`
	Resolution  string   `json:"resolution,omitempty"`
}
