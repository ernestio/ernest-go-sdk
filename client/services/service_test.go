/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package services

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ernestio/ernest-sdk/config"
	"github.com/ernestio/ernest-sdk/connection"
	"github.com/ernestio/ernest-sdk/models"
	"github.com/stretchr/testify/suite"
)

// ServicesTestSuite : Test suite for services
type ServicesTestSuite struct {
	suite.Suite
	Services *Services
}

// SetupTest : sets up test suite
func (suite *ServicesTestSuite) SetupTest() {
	mux := http.NewServeMux()
	mux.HandleFunc(apiroute, testhandler)
	mux.HandleFunc(apiroute+"test", testhandler)
	server := httptest.NewServer(mux)

	conn := connection.New(config.New(server.URL))
	suite.Services = &Services{Conn: conn}
}

func (suite *ServicesTestSuite) TestGet() {
	service, err := suite.Services.Get("test")

	suite.Nil(err)
	suite.Equal(service.ID, "1")
	suite.Equal(service.Name, "test")
}

func (suite *ServicesTestSuite) TestList() {
	services, err := suite.Services.List()

	suite.Nil(err)
	suite.Equal(len(services), 2)
	suite.Equal(services[0].ID, "1")
	suite.Equal(services[0].Name, "test")
	suite.Equal(services[1].ID, "2")
	suite.Equal(services[1].Name, "example")
}

func (suite *ServicesTestSuite) TestCreate() {
	m := &models.Service{
		Name: "test",
	}

	err := suite.Services.Create(m)

	suite.Nil(err)
	suite.Equal(m.ID, "1")
	suite.Equal(m.Name, "test")
}

func (suite *ServicesTestSuite) TestUpdate() {
	m := &models.Service{
		Name: "test",
		Sync: true,
	}

	err := suite.Services.Update(m)

	suite.Nil(err)
	suite.Equal(m.Name, "test")
	suite.Equal(m.Sync, true)
}

func (suite *ServicesTestSuite) TestDelete() {
	service, err := suite.Services.Delete("test")

	suite.Nil(err)
	suite.Equal(service.Name, "test")
}

// TestServicesTestSuite : Test suite for connection
func TestServicesTestSuite(t *testing.T) {
	suite.Run(t, new(ServicesTestSuite))
}
