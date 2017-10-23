/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package loggers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ernestio/ernest-go-sdk/config"
	"github.com/ernestio/ernest-go-sdk/connection"
	"github.com/ernestio/ernest-go-sdk/models"
	"github.com/stretchr/testify/suite"
)

// LoggersTestSuite : Test suite for loggers
type LoggersTestSuite struct {
	suite.Suite
	Loggers *Loggers
}

// SetupTest : sets up test suite
func (suite *LoggersTestSuite) SetupTest() {
	mux := http.NewServeMux()
	mux.HandleFunc(apiroute, testhandler)
	mux.HandleFunc(apiroute+"test", testhandler)
	server := httptest.NewServer(mux)

	conn := connection.New(config.New(server.URL))
	suite.Loggers = &Loggers{Conn: conn}
}

func (suite *LoggersTestSuite) TestList() {
	loggers, err := suite.Loggers.List()

	suite.Nil(err)
	suite.Equal(len(loggers), 1)
	suite.Equal(loggers[0].Type, "tp-1")
	suite.Equal(loggers[0].Logfile, "test-1")
	suite.Equal(loggers[0].Hostname, "host-1")
	suite.Equal(loggers[0].Port, 80)
	suite.Equal(loggers[0].Timeout, 5)
	suite.Equal(loggers[0].Token, "tk-1")
	suite.Equal(loggers[0].Environment, "env-1")
	suite.Equal(loggers[0].UUID, "uuid-1")
}

func (suite *LoggersTestSuite) TestDelete() {
	err := suite.Loggers.Delete("test")

	suite.Nil(err)
}

func (suite *LoggersTestSuite) TestCreate() {
	m := &models.Logger{
		Type:     "tp-1",
		Hostname: "host-1",
	}

	err := suite.Loggers.Create(m)

	suite.Nil(err)
	suite.Equal(m.Type, "tp-1")
	suite.Equal(m.Hostname, "host-1")
}

// TestLoggersTestSuite : Test suite for connection
func TestLoggersTestSuite(t *testing.T) {
	suite.Run(t, new(LoggersTestSuite))
}
