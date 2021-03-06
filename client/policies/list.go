/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package policies

import (
	"github.com/ernestio/ernest-go-sdk/client/generic"
	"github.com/ernestio/ernest-go-sdk/models"
)

// List : list all availabile policies
func (p *Policies) List() (ms []*models.Policy, err error) {
	err = generic.New(p.Conn, apiroute).List(&ms)
	return
}
