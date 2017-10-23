/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package projects

import (
	"fmt"
)

// Delete : delete a project
func (p *Projects) Delete(name string) error {
	path := fmt.Sprintf(apiroute+"%s", name)

	resp, err := p.Conn.Delete(path)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}
