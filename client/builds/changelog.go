/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package builds

import (
	"fmt"

	"github.com/ernestio/ernest-go-sdk/connection"
	"github.com/r3labs/diff"
)

// Changelog : get a environment's build changelog
func (b *Builds) Changelog(project, environment, id string) (*diff.Changelog, error) {
	var m diff.Changelog

	path := fmt.Sprintf(apiroute+"%s/mapping/?changelog=true", project, environment, id)

	resp, err := b.Conn.Get(path)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return &m, connection.ReadJSON(resp.Body, &m)
}
