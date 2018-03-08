/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package environments

import (
	"encoding/json"
	"fmt"

	"github.com/ernestio/ernest-go-sdk/connection"
	"github.com/ernestio/ernest-go-sdk/models"
)

// Validate : validates an environment against its policy documents
func (e *Environments) Validate(project, environment string) (*models.Validation, error) {
	var v models.Validation

	m := models.Action{
		Type: "validate",
	}

	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(apiroute+"%s/actions/", project, environment)
	resp, err := e.Conn.Post(path, "application/json", data)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return &v, connection.ReadJSON(resp.Body, &v)
}
