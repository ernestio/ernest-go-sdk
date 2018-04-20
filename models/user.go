/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

// User : stores user data
type User struct {
	ID                 int    `json:"id"`
	GroupID            int    `json:"group_id"`
	Email              string `json:"email"`
	Username           string `json:"username"`
	Password           string `json:"password,omitempty"`
	OldPassword        string `json:"oldpassword,omitempty"`
	Salt               string `json:"salt,omitempty"`
	Admin              bool   `json:"admin"`
	Type               string `json:"type"`
	MFA                *bool  `json:"mfa"`
	MFASecret          string `json:"mfa_secret"`
	EnvMemberships     []Role `json:"env_memberships"`
	ProjectMemberships []Role `json:"project_memberships"`
	Disabled           *bool  `json:"disabled"`
}

// GetID : get the id for the current object
func (n *User) GetID() string {
	return n.Username
}
