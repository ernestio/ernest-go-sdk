/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package notifications

// RemoveProject : removes a project to a notification
func (u *Notifications) RemoveProject(notification, project string) error {
	route := apiroute + notification + "/projects/" + project
	_, err := u.Conn.Delete(route)

	return err
}
