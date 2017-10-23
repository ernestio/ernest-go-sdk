/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package environments

import "github.com/ernestio/ernest-go-sdk/models"

// Sync : sync's an envionrments data with state stored on the provider
func (e *Environments) Sync(project, environment string) (*models.Action, error) {
	m := models.Action{
		Type: "sync",
	}

	return &m, e.Action(project, environment, &m)
}
