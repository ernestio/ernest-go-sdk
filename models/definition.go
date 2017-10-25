/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import yaml "gopkg.in/yaml.v2"

// Definition : stores basic definition data
type Definition struct {
	Name    string `yaml:"name"`
	Project string `yaml:"project"`
}

// GetID : get the id for the current object
func (d *Definition) GetID() string {
	return d.Name
}

// Load : loads a yaml definition
func (d *Definition) Load(data []byte) error {
	return yaml.Unmarshal(data, d)
}
