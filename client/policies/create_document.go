/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package policies

import (
	"encoding/json"
	"fmt"

	"github.com/ernestio/ernest-go-sdk/connection"
	"github.com/ernestio/ernest-go-sdk/models"
)

// Create : create a policy document revision
func (p *Policies) CreateDocument(policy, definition string) (*models.PolicyDocument, error) {
	m := models.PolicyDocument{
		Definition: definition,
	}

	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(apiroute+"%s/revisions/", policy)

	resp, err := p.Conn.Post(path, "application/yaml", data)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return &m, connection.ReadJSON(resp.Body, &m)
}
