/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package environments

import (
	"fmt"
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
	mux.HandleFunc(fmt.Sprintf(apiroute, "test"), testhandler)
	mux.HandleFunc(fmt.Sprintf(apiroute+"%s", "test", "test"), testhandler)
	mux.HandleFunc(fmt.Sprintf(apiroute+"%s/actions/", "test", "test"), testhandler)
	server := httptest.NewServer(mux)

	conn := connection.New(config.New(server.URL))
	suite.Environments = &Environments{Conn: conn}
}

func (suite *EnvironmentsTestSuite) TestGet() {
	environment, err := suite.Environments.Get("test", "test")

	suite.Nil(err)
	suite.Equal(environment.ID, 1)
	suite.Equal(environment.Name, "test/test")
}

func (suite *EnvironmentsTestSuite) TestList() {
	environments, err := suite.Environments.List("test")

	suite.Nil(err)
	suite.Equal(len(environments), 2)
	suite.Equal(environments[0].ID, 1)
	suite.Equal(environments[0].Name, "test/test")
	suite.Equal(environments[1].ID, 2)
	suite.Equal(environments[1].Name, "test/example")
}

func (suite *EnvironmentsTestSuite) TestCreate() {
	m := &models.Environment{
		Name: "test",
	}

	err := suite.Environments.Create("test", m)

	suite.Nil(err)
	suite.Equal(m.ID, 1)
	suite.Equal(m.Name, "test/test")
}

func (suite *EnvironmentsTestSuite) TestUpdate() {
	m := &models.Environment{
		Name: "test/test",
	}

	err := suite.Environments.Update(m)

	suite.Nil(err)
	suite.Equal(m.Name, "test/test")
}

func (suite *EnvironmentsTestSuite) TestDelete() {
	build, err := suite.Environments.Delete("test", "test")

	suite.Nil(err)
	suite.Equal(build.ID, "1")
	suite.Equal(build.Type, "delete")
	suite.Equal(build.Status, "running")
}

func (suite *EnvironmentsTestSuite) TestAction() {
	action := models.Action{Type: "test"}
	err := suite.Environments.Action("test", "test", &action)

	suite.Nil(err)
	suite.Equal(action.Type, "test")
	suite.Equal(action.Status, "done")
}

func (suite *EnvironmentsTestSuite) TestImport() {
	action, err := suite.Environments.Import("test", "test", []string{"test"})

	suite.Nil(err)
	suite.Equal(action.Type, "import")
	suite.Equal(action.Status, "in_progress")
	suite.Equal(action.ResourceID, "test")
	suite.Equal(action.ResourceType, "build")
	suite.Equal(action.Options.Filters, []string{"test"})
}

func (suite *EnvironmentsTestSuite) TestReset() {
	action, err := suite.Environments.Reset("test", "test")

	suite.Nil(err)
	suite.Equal(action.Type, "reset")
	suite.Equal(action.Status, "done")
}

func (suite *EnvironmentsTestSuite) TestSync() {
	action, err := suite.Environments.Sync("test", "test")

	suite.Nil(err)
	suite.Equal(action.Type, "sync")
	suite.Equal(action.Status, "in_progress")
	suite.Equal(action.ResourceID, "test")
	suite.Equal(action.ResourceType, "build")
}

func (suite *EnvironmentsTestSuite) TestResolve() {
	action, err := suite.Environments.Resolve("test", "test", "reject-changes")

	suite.Nil(err)
	suite.Equal(action.Type, "resolve")
	suite.Equal(action.Options.Resolution, "reject-changes")
	suite.Equal(action.Status, "in_progress")
	suite.Equal(action.ResourceID, "test")
	suite.Equal(action.ResourceType, "build")
}

// TestEnvironmentsTestSuite : Test suite for connection
func TestEnvironmentsTestSuite(t *testing.T) {
	suite.Run(t, new(EnvironmentsTestSuite))
}
