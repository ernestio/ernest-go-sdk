/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package datacenters

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ernestio/ernest-go-sdk/config"
	"github.com/ernestio/ernest-go-sdk/connection"
	"github.com/ernestio/ernest-go-sdk/models"
	"github.com/stretchr/testify/suite"
)

// DatacentersTestSuite : Test suite for datacenters
type DatacentersTestSuite struct {
	suite.Suite
	Datacenters *Datacenters
}

// SetupTest : sets up test suite
func (suite *DatacentersTestSuite) SetupTest() {
	mux := http.NewServeMux()
	mux.HandleFunc(apiroute, testhandler)
	mux.HandleFunc(apiroute+"test", testhandler)
	server := httptest.NewServer(mux)

	conn := connection.New(config.New(server.URL))
	suite.Datacenters = &Datacenters{Conn: conn}
}

func (suite *DatacentersTestSuite) TestGet() {
	datacenter, err := suite.Datacenters.Get(1)

	suite.Nil(err)
	suite.Equal(datacenter.ID, 1)
	suite.Equal(datacenter.Name, "test-1")
}

func (suite *DatacentersTestSuite) TestList() {
	datacenters, err := suite.Datacenters.List()

	suite.Nil(err)
	suite.Equal(len(datacenters), 2)
	suite.Equal(datacenters[0].ID, 1)
	suite.Equal(datacenters[0].Name, "test-1")
	suite.Equal(datacenters[1].ID, 2)
	suite.Equal(datacenters[1].Name, "test-2")
}

func (suite *DatacentersTestSuite) TestCreate() {
	m := &models.Datacenter{
		ID:       1,
		Name:     "test",
		Username: "test",
		Password: "test",
	}

	err := suite.Datacenters.Create(m)

	suite.Nil(err)
	suite.Equal(m.ID, 1)
	suite.Equal(m.Name, "test")
	suite.Equal(m.Username, "test")
	suite.Equal(m.Password, "test")
}

func (suite *DatacentersTestSuite) TestUpdate() {
	m := &models.Datacenter{
		ID:   1,
		Name: "test-1",
	}

	err := suite.Datacenters.Update(m)

	suite.Nil(err)
	suite.Equal(m.Name, "test-1")
}

func (suite *DatacentersTestSuite) TestDelete() {
	err := suite.Datacenters.Delete(1)

	suite.Nil(err)
}

// TestDatacentersTestSuite : Test suite for connection
func TestDatacentersTestSuite(t *testing.T) {
	suite.Run(t, new(DatacentersTestSuite))
}
