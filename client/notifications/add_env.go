/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package notifications

// AddEnv : adds an environment to a notification
func (u *Notifications) AddEnv(notification, project, env string) error {
	route := apiroute + notification + "/projects/" + project + "/envs/" + env
	_, err := u.Conn.Post(route, "application/json", []byte(""))

	return err
}
