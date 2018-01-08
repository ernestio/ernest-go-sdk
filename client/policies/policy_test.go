/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package policies

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ernestio/ernest-go-sdk/config"
	"github.com/ernestio/ernest-go-sdk/connection"
	"github.com/ernestio/ernest-go-sdk/models"
	"github.com/stretchr/testify/suite"
)

// PoliciesTestSuite : Test suite for policies
type PoliciesTestSuite struct {
	suite.Suite
	Policies *Policies
}

// SetupTest : sets up test suite
func (suite *PoliciesTestSuite) SetupTest() {
	mux := http.NewServeMux()
	mux.HandleFunc(apiroute, testhandler)
	mux.HandleFunc(apiroute+"test", testhandler)
	server := httptest.NewServer(mux)

	conn := connection.New(config.New(server.URL))
	suite.Policies = &Policies{Conn: conn}
}

func (suite *PoliciesTestSuite) TestGet() {
	notification, err := suite.Policies.Get("test-1")

	suite.Nil(err)
	suite.Equal(notification.ID, 1)
	suite.Equal(notification.Name, "test-1")
	suite.Equal(notification.Definition, "type-1")
}

func (suite *PoliciesTestSuite) TestList() {
	policies, err := suite.Policies.List()

	suite.Nil(err)
	suite.Equal(len(policies), 2)
	suite.Equal(policies[0].ID, 1)
	suite.Equal(policies[0].Name, "test-1")
	suite.Equal(policies[0].Definition, "type-1")
	suite.Equal(policies[1].ID, 2)
	suite.Equal(policies[1].Name, "test-2")
	suite.Equal(policies[1].Definition, "type-2")
}

func (suite *PoliciesTestSuite) TestCreate() {
	notification := &models.Policy{
		ID:         1,
		Name:       "test-1",
		Definition: "type-1",
	}

	err := suite.Policies.Create(notification)

	suite.Nil(err)
	suite.Equal(notification.ID, 1)
	suite.Equal(notification.Name, "test-1")
	suite.Equal(notification.Definition, "type-1")
}

func (suite *PoliciesTestSuite) TestUpdate() {
	notification := &models.Policy{
		ID:         1,
		Name:       "test-1",
		Definition: "type-1",
	}

	err := suite.Policies.Update(notification)

	suite.Nil(err)
	suite.Equal(notification.ID, 1)
	suite.Equal(notification.Name, "test-1")
	suite.Equal(notification.Definition, "type-1")
}

func (suite *PoliciesTestSuite) TestDelete() {
	err := suite.Policies.Delete("test-1")

	suite.Nil(err)
}

// TestPoliciesTestSuite : Test suite for connection
func TestPoliciesTestSuite(t *testing.T) {
	suite.Run(t, new(PoliciesTestSuite))
}
