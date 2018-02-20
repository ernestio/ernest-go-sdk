/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import (
	"strings"
)

// Validation describes a response from the build validate service.
type Validation struct {
	Version    string     `json:"version"`
	Controls   []Control  `json:"controls"`
	Profiles   []Profile  `json:"profiles"`
	Statistics Statistics `json:"statistics"`
}

// Profile describes the policy document and its test results
type Profile struct {
	Supports []string         `json:"supports"`
	Title    string           `json:"title"`
	Name     string           `json:"name"`
	Controls []ControlDetails `json:"controls"`
}

// ControlDetails describes additional information about a control
type ControlDetails struct {
	ID             string    `json:"id"`
	Title          string    `json:"title"`
	Description    string    `json:"desc"`
	Impact         float32   `json:"impact"`
	References     []string  `json:"refs"`
	Tags           []string  `json:"tags"`
	Code           string    `json:"code"`
	Results        []Control `json:"results"`
	Groups         []Group   `json:"groups"`
	Attributes     []string  `json:"attribues"`
	SHA256         string    `json:"sha256"`
	SourceLocation struct {
		Reference string `json:"ref"`
		Line      int    `json:"line"`
	} `json:"source_location"`
}

// Control describes an individual test within a build validation.
type Control struct {
	ID        string  `json:"id"`
	ProfileID string  `json:"profile_id"`
	Status    string  `json:"status"`
	CodeDesc  string  `json:"code_desc"`
	RunTime   float64 `json:"run_time"`
	StartTime string  `json:"start_time"`
	Message   string  `json:"message"`
}

type Group struct {
	ID       string   `json:"id"`
	Title    string   `json:"title"`
	Controls []string `json:"controls"`
}

// Statistics describes stats for the build validate service.
type Statistics struct {
	Duration float64 `json:"duration"`
}

// Passed : returns true if validation rules passed
func (v *Validation) Passed() bool {
	for _, c := range v.Controls {
		if c.Status == "failed" {
			return false
		}
	}

	return true
}

func (v *Validation) Stats() (int, int, int) {
	var passed, failed int

	total := len(v.Controls)

	for _, c := range v.Controls {
		if c.Status == "passed" {
			passed++
		} else {
			failed++
		}
	}

	return passed, failed, total
}

// PolicyName : Returns the name of the policy that the control is derrived from
func (c *Control) PolicyName() string {
	values := strings.Split(c.ID, " ")
	pn := strings.Split(values[2], ".rb")
	// remove 37 additional characters (uuid plus dashes)
	return pn[0][:(len(pn[0]) - 37)]
}

// Line : The line position on the policy that the control references
func (c *Control) Line() string {
	values := strings.Split(c.ID, " ")
	pn := strings.Split(values[2], ".rb")
	return pn[1][1:]
}
