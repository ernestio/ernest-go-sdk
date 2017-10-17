/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package client

import (
	"github.com/ernestio/ernest-go-sdk/client/builds"
	"github.com/ernestio/ernest-go-sdk/client/environments"
	"github.com/ernestio/ernest-go-sdk/client/groups"
	"github.com/ernestio/ernest-go-sdk/client/projects"
	"github.com/ernestio/ernest-go-sdk/client/users"
	"github.com/ernestio/ernest-go-sdk/config"
	"github.com/ernestio/ernest-go-sdk/connection"
)

// Client :
type Client struct {
	Conn         *connection.Conn
	Environments *environments.Environments
	Builds       *builds.Builds
	Projects     *projects.Projects
	Groups       *groups.Groups
	Users        *users.Users
}

// New : creates a new client
func New(cfg *config.Config) *Client {
	c := connection.New(cfg)

	return &Client{
		Conn:         c,
		Environments: &environments.Environments{Conn: c},
		Builds:       &builds.Builds{Conn: c},
		Projects:     &projects.Projects{Conn: c},
		Groups:       &groups.Groups{Conn: c},
		Users:        &users.Users{Conn: c},
	}
}
