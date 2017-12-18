/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package client

import (
	"github.com/ernestio/ernest-go-sdk/client/builds"
	"github.com/ernestio/ernest-go-sdk/client/environments"
	"github.com/ernestio/ernest-go-sdk/client/loggers"
	"github.com/ernestio/ernest-go-sdk/client/notifications"
	"github.com/ernestio/ernest-go-sdk/client/projects"
	"github.com/ernestio/ernest-go-sdk/client/roles"
	"github.com/ernestio/ernest-go-sdk/client/sessions"
	"github.com/ernestio/ernest-go-sdk/client/users"
	"github.com/ernestio/ernest-go-sdk/config"
	"github.com/ernestio/ernest-go-sdk/connection"
)

// Client :
type Client struct {
	Conn          *connection.Conn
	Environments  *environments.Environments
	Builds        *builds.Builds
	Projects      *projects.Projects
	Users         *users.Users
	Sessions      *sessions.Sessions
	Notifications *notifications.Notifications
	Roles         *roles.Roles
	Loggers       *loggers.Loggers
}

// New : creates a new client
func New(cfg *config.Config) *Client {
	c := connection.New(cfg)

	return &Client{
		Conn:          c,
		Environments:  &environments.Environments{Conn: c},
		Builds:        &builds.Builds{Conn: c},
		Projects:      &projects.Projects{Conn: c},
		Users:         &users.Users{Conn: c},
		Sessions:      &sessions.Sessions{Conn: c},
		Notifications: &notifications.Notifications{Conn: c},
		Roles:         &roles.Roles{Conn: c},
		Loggers:       &loggers.Loggers{Conn: c},
	}
}
