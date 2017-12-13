/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package roles

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ernestio/ernest-go-sdk/config"
	"github.com/ernestio/ernest-go-sdk/connection"
	"github.com/ernestio/ernest-go-sdk/models"
	"github.com/stretchr/testify/suite"
)

// RolesTestSuite : Test suite for roles
type RolesTestSuite struct {
	suite.Suite
	Roles *Roles
}

// SetupTest : sets up test suite
func (suite *RolesTestSuite) SetupTest() {
	mux := http.NewServeMux()
	mux.HandleFunc(apiroute, testhandler)
	mux.HandleFunc(apiroute+"test", testhandler)
	server := httptest.NewServer(mux)

	conn := connection.New(config.New(server.URL))
	suite.Roles = &Roles{Conn: conn}
}

func (suite *RolesTestSuite) TestCreate() {
	role := &models.Role{
		ID:       "test-1",
		User:     "usr-1",
		Role:     "rol-1",
		Resource: "res-1",
	}

	err := suite.Roles.Create(role)

	suite.Nil(err)
	suite.Equal(role.ID, "test-1")
	suite.Equal(role.User, "usr-1")
	suite.Equal(role.Role, "rol-1")
	suite.Equal(role.Resource, "res-1")
}

func (suite *RolesTestSuite) TestDelete() {
	role := &models.Role{
		ID:       "test-1",
		User:     "usr-1",
		Role:     "rol-1",
		Resource: "res-1",
	}

	err := suite.Roles.Delete(role)
	suite.Nil(err)
}

// TestRolesTestSuite : Test suite for connection
func TestRolesTestSuite(t *testing.T) {
	suite.Run(t, new(RolesTestSuite))
}
