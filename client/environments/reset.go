/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package environments

import "github.com/ernestio/ernest-go-sdk/models"

// Reset : resets an environments state
func (e *Environments) Reset(project, environment string) (*models.Action, error) {
	m := models.Action{
		Type: "reset",
	}

	return &m, e.Action(project, environment, &m)
}
