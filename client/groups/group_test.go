/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package groups

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ernestio/ernest-sdk/config"
	"github.com/ernestio/ernest-sdk/connection"
	"github.com/ernestio/ernest-sdk/models"
	"github.com/stretchr/testify/suite"
)

// GroupsTestSuite : Test suite for groups
type GroupsTestSuite struct {
	suite.Suite
	Groups *Groups
}

// SetupTest : sets up test suite
func (suite *GroupsTestSuite) SetupTest() {
	mux := http.NewServeMux()
	mux.HandleFunc(apiroute, testhandler)
	mux.HandleFunc(apiroute+"test", testhandler)
	server := httptest.NewServer(mux)

	conn := connection.New(config.New(server.URL))
	suite.Groups = &Groups{Conn: conn}
}

func (suite *GroupsTestSuite) TestGet() {
	group, err := suite.Groups.Get(1)

	suite.Nil(err)
	suite.Equal(group.ID, 1)
	suite.Equal(group.Name, "test-1")
}

func (suite *GroupsTestSuite) TestList() {
	groups, err := suite.Groups.List()

	suite.Nil(err)
	suite.Equal(len(groups), 2)
	suite.Equal(groups[0].ID, 1)
	suite.Equal(groups[0].Name, "test-1")
	suite.Equal(groups[1].ID, 2)
	suite.Equal(groups[1].Name, "test-2")
}

func (suite *GroupsTestSuite) TestCreate() {
	m := &models.Group{
		Name: "test",
	}

	err := suite.Groups.Create(m)

	suite.Nil(err)
	suite.Equal(m.ID, 1)
	suite.Equal(m.Name, "test")
}

func (suite *GroupsTestSuite) TestDelete() {
	err := suite.Groups.Delete(1)

	suite.Nil(err)
}

// TestGroupsTestSuite : Test suite for connection
func TestGroupsTestSuite(t *testing.T) {
	suite.Run(t, new(GroupsTestSuite))
}
