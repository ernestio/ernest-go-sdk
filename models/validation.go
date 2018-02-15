/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

// Validation describes a response from the build validate service.
type Validation struct {
	Version    string     `json:"version"`
	Controls   []Control  `json:"controls"`
	Statistics Statistics `json:"statistics"`
}

// Control describes an individual test within a build validation.
type Control struct {
	ID        string `json:"id"`
	ProfileID string `json:"profile_id"`
	Status    string `json:"status"`
	CodeDesc  string `json:"code_desc"`
	Message   string `json:"message"`
}

// Statistics describes stats for the build validate service.
type Statistics struct {
	Duration float64 `json:"duration"`
}

// Passed : returns true if validation rules passed
func (b *Validation) Passed() bool {
	for _, e := range b.Controls {
		if e.Status == "failed" {
			return false
		}
	}

	return true
}
