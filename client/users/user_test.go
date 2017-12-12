/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package users

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ernestio/ernest-go-sdk/config"
	"github.com/ernestio/ernest-go-sdk/connection"
	"github.com/ernestio/ernest-go-sdk/models"
	"github.com/stretchr/testify/suite"
)

// UsersTestSuite : Test suite for users
type UsersTestSuite struct {
	suite.Suite
	Users *Users
}

// SetupTest : sets up test suite
func (suite *UsersTestSuite) SetupTest() {
	mux := http.NewServeMux()
	mux.HandleFunc(apiroute, testhandler)
	mux.HandleFunc(apiroute+"test", testhandler)
	server := httptest.NewServer(mux)

	conn := connection.New(config.New(server.URL))
	suite.Users = &Users{Conn: conn}
}

func (suite *UsersTestSuite) TestGet() {
	user, err := suite.Users.Get("1")

	suite.Nil(err)
	suite.Equal(user.ID, 1)
	suite.Equal(user.Username, "test-1")
}

func (suite *UsersTestSuite) TestList() {
	users, err := suite.Users.List()

	suite.Nil(err)
	suite.Equal(len(users), 2)
	suite.Equal(users[0].ID, 1)
	suite.Equal(users[0].Username, "test-1")
	suite.Equal(users[1].ID, 2)
	suite.Equal(users[1].Username, "test-2")
}

func (suite *UsersTestSuite) TestCreate() {
	m := &models.User{
		ID:       1,
		Username: "test",
	}

	err := suite.Users.Create(m)

	suite.Nil(err)
	suite.Equal(m.ID, 1)
	suite.Equal(m.Username, "test")
}

func (suite *UsersTestSuite) TestUpdate() {
	m := &models.User{
		ID:       1,
		Username: "test-1",
	}

	err := suite.Users.Update(m)

	suite.Nil(err)
	suite.Equal(m.Username, "test-1")
}

func (suite *UsersTestSuite) TestDelete() {
	err := suite.Users.Delete(1)

	suite.Nil(err)
}

// TestUsersTestSuite : Test suite for connection
func TestUsersTestSuite(t *testing.T) {
	suite.Run(t, new(UsersTestSuite))
}
