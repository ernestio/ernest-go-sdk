/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package environments

import "github.com/ernestio/ernest-go-sdk/models"

// Import : creates a an import build for a environment
func (e *Environments) Import(project, environment string, filters []string) (*models.Action, error) {
	m := models.Action{
		Type: "import",
		Options: models.ActionOptions{
			Filters: filters,
		},
	}

	return &m, e.Action(project, environment, &m)
}
