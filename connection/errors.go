/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package connection

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// ErnestError : ...
type ErnestError interface {
	Error() string
	Status() int
}

type ernestError struct {
	status          int
	message         string
	originalMessage string
}

func newError(status int, message string) ErnestError {
	return &ernestError{
		status:  status,
		message: message,
	}
}

// Error ...
func (e ernestError) Error() string {
	return e.message
}

// Status ...
func (e ernestError) Status() int {
	return e.status
}

func responseError(resp *http.Response) ErnestError {
	var err *ernestError

	stErr := status(resp.StatusCode)
	if stErr == nil {
		return nil
	}
	err = stErr.(*ernestError)

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		log.Println(err.Error())
		return err
	}

	var respMsg struct {
		Message string `json:"message"`
	}

	if err2 := json.Unmarshal(body, &respMsg); err2 == nil {
		err.message = respMsg.Message
	}

	return err
}

func status(s int) ErnestError {
	switch s {
	case http.StatusOK, http.StatusCreated, http.StatusAccepted:
		return nil
	case http.StatusBadRequest:
		return newError(s, "ernest responded with code 400 : 'bad request'")
	case http.StatusUnauthorized:
		return newError(s, "you are not authorized to perform this action, please log in")
	case http.StatusForbidden:
		return newError(s, "you are not autorized to perform this action with your level of permissions")
	case http.StatusNotFound:
		return newError(s, "the resource does not exist")
	case http.StatusInternalServerError:
		return newError(s, "ernest responded with code 500 : 'internal server error'")
	default:
		return newError(s, "ernest unknown error")
	}
}
