/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package notifications

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ernestio/ernest-go-sdk/config"
	"github.com/ernestio/ernest-go-sdk/connection"
	"github.com/ernestio/ernest-go-sdk/models"
	"github.com/stretchr/testify/suite"
)

// NotificationsTestSuite : Test suite for notifications
type NotificationsTestSuite struct {
	suite.Suite
	Notifications *Notifications
}

// SetupTest : sets up test suite
func (suite *NotificationsTestSuite) SetupTest() {
	mux := http.NewServeMux()
	mux.HandleFunc(apiroute, testhandler)
	mux.HandleFunc(apiroute+"test", testhandler)
	server := httptest.NewServer(mux)

	conn := connection.New(config.New(server.URL))
	suite.Notifications = &Notifications{Conn: conn}
}

func (suite *NotificationsTestSuite) TestGet() {
	notification, err := suite.Notifications.Get("test-1")

	suite.Nil(err)
	suite.Equal(notification.ID, 1)
	suite.Equal(notification.Name, "test-1")
	suite.Equal(notification.Type, "type-1")
	suite.Equal(notification.Config, "config-1")
	suite.Equal(notification.Sources[0], "a1")
	suite.Equal(notification.Sources[1], "b1")
}

func (suite *NotificationsTestSuite) TestList() {
	notifications, err := suite.Notifications.List()

	suite.Nil(err)
	suite.Equal(len(notifications), 2)
	suite.Equal(notifications[0].ID, 1)
	suite.Equal(notifications[0].Name, "test-1")
	suite.Equal(notifications[0].Type, "type-1")
	suite.Equal(notifications[0].Config, "config-1")
	suite.Equal(notifications[0].Sources[0], "a1")
	suite.Equal(notifications[0].Sources[1], "b1")

	suite.Equal(notifications[1].ID, 2)
	suite.Equal(notifications[1].Name, "test-2")
	suite.Equal(notifications[1].Type, "type-2")
	suite.Equal(notifications[1].Config, "config-2")
	suite.Equal(notifications[1].Sources[0], "a2")
	suite.Equal(notifications[1].Sources[1], "b2")
}

func (suite *NotificationsTestSuite) TestCreate() {
	notification := &models.Notification{
		ID:      1,
		Name:    "test-1",
		Type:    "type-1",
		Config:  "config-1",
		Sources: []string{"a1", "b1"},
	}

	err := suite.Notifications.Create(notification)

	suite.Nil(err)
	suite.Equal(notification.ID, 1)
	suite.Equal(notification.Name, "test-1")
	suite.Equal(notification.Type, "type-1")
	suite.Equal(notification.Config, "config-1")
	suite.Equal(notification.Sources[0], "a1")
	suite.Equal(notification.Sources[1], "b1")
}

func (suite *NotificationsTestSuite) TestUpdate() {
	notification := &models.Notification{
		ID:      1,
		Name:    "test-1",
		Type:    "type-1",
		Config:  "config-1",
		Sources: []string{"a1", "b1"},
	}

	err := suite.Notifications.Update(notification)

	suite.Nil(err)
	suite.Equal(notification.ID, 1)
	suite.Equal(notification.Name, "test-1")
	suite.Equal(notification.Type, "type-1")
	suite.Equal(notification.Config, "config-1")
	suite.Equal(notification.Sources[0], "a1")
	suite.Equal(notification.Sources[1], "b1")
}

func (suite *NotificationsTestSuite) TestDelete() {
	err := suite.Notifications.Delete("test-1")

	suite.Nil(err)
}

// TestNotificationsTestSuite : Test suite for connection
func TestNotificationsTestSuite(t *testing.T) {
	suite.Run(t, new(NotificationsTestSuite))
}
