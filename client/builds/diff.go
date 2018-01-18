/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package builds

import (
	"encoding/json"
	"fmt"

	"github.com/ernestio/ernest-go-sdk/connection"
	"github.com/ernestio/ernest-go-sdk/models"
	"github.com/r3labs/diff"
)

// Diff : simulates a build creation
func (b *Builds) Diff(project, environment, from, to string) (*diff.Changelog, error) {
	var m diff.Changelog

	dr := models.Diff{
		From: from,
		To:   to,
	}

	data, err := json.Marshal(dr)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("/api/projects/%s/envs/%s/diff/", project, environment)

	resp, err := b.Conn.Post(path, "application/json", data)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return &m, connection.ReadJSON(resp.Body, &m)
}
