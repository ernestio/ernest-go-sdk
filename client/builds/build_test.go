/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package builds

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ernestio/ernest-sdk/config"
	"github.com/ernestio/ernest-sdk/connection"
	"github.com/ernestio/ernest-sdk/models"
	"github.com/stretchr/testify/suite"
)

// BuildsTestSuite : Test suite for builds
type BuildsTestSuite struct {
	suite.Suite
	Builds *Builds
}

// SetupTest : sets up test suite
func (suite *BuildsTestSuite) SetupTest() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/services/test/import/", testhandler)
	mux.HandleFunc("/api/services/test/builds/", testhandler)
	mux.HandleFunc("/api/services/test/builds/test", testhandler)
	server := httptest.NewServer(mux)

	conn := connection.New(config.New(server.URL))
	suite.Builds = &Builds{Conn: conn}
}

func (suite *BuildsTestSuite) TestGet() {
	build, err := suite.Builds.Get("test", "1")

	suite.Nil(err)
	suite.Equal(build.ID, "1")
	suite.Equal(build.Status, "completed")
}

func (suite *BuildsTestSuite) TestList() {
	builds, err := suite.Builds.List("test")

	suite.Nil(err)
	suite.Equal(len(builds), 2)
	suite.Equal(builds[0].ID, "1")
	suite.Equal(builds[0].Status, "completed")
	suite.Equal(builds[1].ID, "2")
	suite.Equal(builds[1].Status, "running")
}

func (suite *BuildsTestSuite) TestCreate() {
	data := []byte(`
		---
		name: test
		datacenter: datacenter
	`)

	build, err := suite.Builds.Create("test", data)

	suite.Nil(err)
	suite.Equal(build.ID, "1")
	suite.Equal(build.Status, "running")
}

func (suite *BuildsTestSuite) TestImport() {
	m := &models.Import{
		Filters: []string{"test"},
	}

	build, err := suite.Builds.Import("test", m)

	suite.Nil(err)
	suite.Equal(build.ID, "1")
	suite.Equal(build.Type, "import")
	suite.Equal(build.Status, "running")
}

// TestBuildsTestSuite : Test suite for connection
func TestBuildsTestSuite(t *testing.T) {
	suite.Run(t, new(BuildsTestSuite))
}
