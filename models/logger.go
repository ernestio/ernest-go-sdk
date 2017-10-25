/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

// Logger : stores user data
type Logger struct {
	Type        string `json:"type"`
	Logfile     string `json:"logfile"`
	Hostname    string `json:"hostname"`
	Port        int    `json:"port"`
	Timeout     int    `json:"timeout"`
	Token       string `json:"token"`
	Environment string `json:"environment"`
	UUID        string `json:"uuid"`
}

// GetID : get the id for the current object
func (n *Logger) GetID() string {
	return n.Type
}
