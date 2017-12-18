/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package builds

import (
	"bytes"
	"fmt"
)

// Definition : get a environment's build definition
func (b *Builds) Definition(project, environment, id string) (string, error) {
	path := fmt.Sprintf(apiroute+"%s/definition/", project, environment, id)
	resp, err := b.Conn.Get(path)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	return buf.String(), nil
}
