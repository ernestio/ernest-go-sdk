/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package environments

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ernestio/ernest-go-sdk/config"
	"github.com/ernestio/ernest-go-sdk/connection"
	"github.com/ernestio/ernest-go-sdk/models"
	"github.com/stretchr/testify/suite"
)

// EnvironmentsTestSuite : Test suite for environments
type EnvironmentsTestSuite struct {
	suite.Suite
	Environments *Environments
}

// SetupTest : sets up test suite
func (suite *EnvironmentsTestSuite) SetupTest() {
	mux := http.NewServeMux()
	mux.HandleFunc(apiroute, testhandler)
	mux.HandleFunc(apiroute+"test", testhandler)
	server := httptest.NewServer(mux)

	conn := connection.New(config.New(server.URL))
	suite.Environments = &Environments{Conn: conn}
}

func (suite *EnvironmentsTestSuite) TestGet() {
	environment, err := suite.Environments.Get("test")

	suite.Nil(err)
	suite.Equal(environment.ID, 1)
	suite.Equal(environment.Name, "test")
}

func (suite *EnvironmentsTestSuite) TestList() {
	environments, err := suite.Environments.List()

	suite.Nil(err)
	suite.Equal(len(environments), 2)
	suite.Equal(environments[0].ID, 1)
	suite.Equal(environments[0].Name, "test")
	suite.Equal(environments[1].ID, 2)
	suite.Equal(environments[1].Name, "example")
}

func (suite *EnvironmentsTestSuite) TestCreate() {
	m := &models.Environment{
		Name: "test",
	}

	err := suite.Environments.Create(m)

	suite.Nil(err)
	suite.Equal(m.ID, 1)
	suite.Equal(m.Name, "test")
}

func (suite *EnvironmentsTestSuite) TestUpdate() {
	m := &models.Environment{
		Name: "test",
	}

	err := suite.Environments.Update(m)

	suite.Nil(err)
	suite.Equal(m.Name, "test")
}

func (suite *EnvironmentsTestSuite) TestDelete() {
	build, err := suite.Environments.Delete("test")

	suite.Nil(err)
	suite.Equal(build.ID, "1")
	suite.Equal(build.Type, "delete")
	suite.Equal(build.Status, "running")
}

// TestEnvironmentsTestSuite : Test suite for connection
func TestEnvironmentsTestSuite(t *testing.T) {
	suite.Run(t, new(EnvironmentsTestSuite))
}
