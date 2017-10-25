/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package notifications

import (
	"github.com/ernestio/ernest-go-sdk/client/generic"
	"github.com/ernestio/ernest-go-sdk/models"
)

// Get : get a notification
func (l *Notifications) Get(id string) (m *models.Notification, err error) {
	err = generic.New(l.Conn, apiroute).Get(id, &m)
	return m, err
}
