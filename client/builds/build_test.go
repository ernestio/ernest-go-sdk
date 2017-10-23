/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package builds

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ernestio/ernest-go-sdk/config"
	"github.com/ernestio/ernest-go-sdk/connection"
	"github.com/r3labs/sse"
	"github.com/stretchr/testify/suite"
)

// BuildsTestSuite : Test suite for builds
type BuildsTestSuite struct {
	server *sse.Server
	suite.Suite
	Builds *Builds
}

// SetupTest : sets up test suite
func (suite *BuildsTestSuite) SetupTest() {
	suite.server = sse.New()
	suite.server.CreateStream("test")

	suite.server.Publish("test", &sse.Event{Data: []byte("test-1")})
	suite.server.Publish("test", &sse.Event{Data: []byte("test-2")})

	mux := http.NewServeMux()

	mux.HandleFunc(fmt.Sprintf(apiroute, "test", "test"), testhandler)
	mux.HandleFunc(fmt.Sprintf(apiroute+"%s", "test", "test", "1"), testhandler)
	mux.HandleFunc("/events", suite.server.HTTPHandler)

	server := httptest.NewServer(mux)

	conn := connection.New(config.New(server.URL))
	suite.Builds = &Builds{Conn: conn}
}

func (suite *BuildsTestSuite) TestGet() {
	build, err := suite.Builds.Get("test", "test", "1")

	suite.Nil(err)
	suite.Equal(build.ID, "1")
	suite.Equal(build.Status, "completed")
}

func (suite *BuildsTestSuite) TestList() {
	builds, err := suite.Builds.List("test", "test")

	suite.Nil(err)
	suite.Equal(len(builds), 2)
	suite.Equal(builds[0].ID, "1")
	suite.Equal(builds[0].Status, "completed")
	suite.Equal(builds[1].ID, "2")
	suite.Equal(builds[1].Status, "running")
}

func (suite *BuildsTestSuite) TestCreate() {
	data := []byte("---\nname: test \nproject: test")

	build, err := suite.Builds.Create(data)

	suite.Nil(err)
	suite.Equal(build.ID, "1")
	suite.Equal(build.Status, "running")
}

func (suite *BuildsTestSuite) TestStream() {
	var events []*sse.Event

	stream, err := suite.Builds.Stream("test")
	suite.Nil(err)

	for i := 0; i < 2; i++ {
		e, ok := <-stream
		if !ok {
			break
		}
		if e.Data != nil {
			events = append(events, e)
		} else {
			i--
		}
	}

	suite.Equal(len(events), 2)
	suite.Equal(string(events[0].Data), "test-1")
	suite.Equal(string(events[1].Data), "test-2")
}

// TestBuildsTestSuite : Test suite for connection
func TestBuildsTestSuite(t *testing.T) {
	suite.Run(t, new(BuildsTestSuite))
}
