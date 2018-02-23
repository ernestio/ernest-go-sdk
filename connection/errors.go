/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package connection

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/ernestio/ernest-go-sdk/models"
)

func newError(message string) *models.Error {
	return &models.Error{
		Message: message,
	}
}

func responseError(resp *http.Response) error {
	err := status(resp.StatusCode)
	if err == nil {
		return nil
	}

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		return err
	}

	if json.Unmarshal(body, err) != nil {
		return err
	}

	return err
}

func status(s int) error {
	switch s {
	case http.StatusOK, http.StatusCreated, http.StatusAccepted:
		return nil
	case http.StatusBadRequest:
		return newError("ernest responded with code 400 : 'bad request'")
	case http.StatusUnauthorized:
		return newError("you are not authorized to perform this action, please log in")
	case http.StatusForbidden:
		return newError("you are not autorized to perform this action with your level of permissions")
	case http.StatusNotFound:
		return newError("the resource does not exist")
	case http.StatusInternalServerError:
		return newError("ernest responded with code 500 : 'internal server error'")
	default:
		return newError("ernest unknown error : " + strconv.Itoa(s))
	}
}
