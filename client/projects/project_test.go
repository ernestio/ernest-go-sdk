/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package projects

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ernestio/ernest-go-sdk/config"
	"github.com/ernestio/ernest-go-sdk/connection"
	"github.com/ernestio/ernest-go-sdk/models"
	"github.com/stretchr/testify/suite"
)

// ProjectsTestSuite : Test suite for projects
type ProjectsTestSuite struct {
	suite.Suite
	Projects *Projects
}

// SetupTest : sets up test suite
func (suite *ProjectsTestSuite) SetupTest() {
	mux := http.NewServeMux()
	mux.HandleFunc(apiroute, testhandler)
	mux.HandleFunc(apiroute+"test", testhandler)
	server := httptest.NewServer(mux)

	conn := connection.New(config.New(server.URL))
	suite.Projects = &Projects{Conn: conn}
}

func (suite *ProjectsTestSuite) TestGet() {
	project, err := suite.Projects.Get(1)

	suite.Nil(err)
	suite.Equal(project.ID, 1)
	suite.Equal(project.Name, "test-1")
}

func (suite *ProjectsTestSuite) TestList() {
	projects, err := suite.Projects.List()

	suite.Nil(err)
	suite.Equal(len(projects), 2)
	suite.Equal(projects[0].ID, 1)
	suite.Equal(projects[0].Name, "test-1")
	suite.Equal(projects[1].ID, 2)
	suite.Equal(projects[1].Name, "test-2")
}

func (suite *ProjectsTestSuite) TestCreate() {
	m := &models.Project{
		ID:   1,
		Name: "test",
		Credentials: map[string]interface{}{
			"username": "test",
			"password": "test",
		},
	}

	err := suite.Projects.Create(m)

	suite.Nil(err)
	suite.Equal(m.ID, 1)
	suite.Equal(m.Name, "test")
	suite.Equal(m.Credentials["username"], "test")
	suite.Equal(m.Credentials["password"], "test")
}

func (suite *ProjectsTestSuite) TestUpdate() {
	m := &models.Project{
		ID:   1,
		Name: "test-1",
	}

	err := suite.Projects.Update(m)

	suite.Nil(err)
	suite.Equal(m.Name, "test-1")
}

func (suite *ProjectsTestSuite) TestDelete() {
	err := suite.Projects.Delete(1)

	suite.Nil(err)
}

// TestProjectsTestSuite : Test suite for connection
func TestProjectsTestSuite(t *testing.T) {
	suite.Run(t, new(ProjectsTestSuite))
}
