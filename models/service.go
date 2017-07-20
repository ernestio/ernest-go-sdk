/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

// Service : stores service data
type Service struct {
	ID           int    `json:"id"`
	GroupID      int    `json:"group_id"`
	DatacenterID int    `json:"datacenter_id"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Status       string `json:"status"`
	Sync         bool   `json:"sync"`
	SyncType     string `json:"sync_type"`
	SyncInterval int    `json:"sync_interval"`
}
