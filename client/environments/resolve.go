/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package environments

import "github.com/ernestio/ernest-go-sdk/models"

// Resolve : resolves a blocking issue/state for an environment
func (e *Environments) Resolve(project, environment, resolution string) (*models.Action, error) {
	m := models.Action{
		Type: "resolve",
		Options: models.ActionOptions{
			Resolution: resolution,
		},
	}

	return &m, e.Action(project, environment, &m)
}
