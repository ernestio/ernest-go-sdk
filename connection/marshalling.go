/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package connection

import (
	"encoding/json"
	"io"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// ReadJSON : reads a json response into an interface
func ReadJSON(body io.ReadCloser, x interface{}) error {
	defer body.Close()
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, x)
}

// ReadYAML reads
func ReadYAML(body io.ReadCloser, x interface{}) error {
	defer body.Close()
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, x)
}
