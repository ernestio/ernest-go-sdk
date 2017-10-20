/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package environments

import (
	"fmt"

	"github.com/ernestio/ernest-go-sdk/connection"
	"github.com/ernestio/ernest-go-sdk/models"
)

// Import : creates a an import build for a environment
func (e *Environments) Import(project, environment string, filters []string) (*models.Build, error) {
	var m models.Build

	a := models.Action{
		Type: "import",
	}
	a.Options.Filters = filters

	err := e.Action(project, environment, &a)
	if err != nil {
		return nil, err
	}

	resp, err := e.Conn.Get(fmt.Sprintf(apiroute + "/builds/%s"))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return &m, connection.ReadJSON(resp.Body, &m)
}
