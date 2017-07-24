/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package client

import (
	"github.com/ernestio/ernest-go-sdk/client/builds"
	"github.com/ernestio/ernest-go-sdk/client/datacenters"
	"github.com/ernestio/ernest-go-sdk/client/groups"
	"github.com/ernestio/ernest-go-sdk/client/services"
	"github.com/ernestio/ernest-go-sdk/client/users"
	"github.com/ernestio/ernest-go-sdk/config"
	"github.com/ernestio/ernest-go-sdk/connection"
)

// Client :
type Client struct {
	Conn        *connection.Conn
	Services    *services.Services
	Builds      *builds.Builds
	Datacenters *datacenters.Datacenters
	Groups      *groups.Groups
	Users       *users.Users
}

// New : creates a new client
func New(cfg *config.Config) *Client {
	c := connection.New(cfg)

	return &Client{
		Conn:        c,
		Services:    &services.Services{Conn: c},
		Builds:      &builds.Builds{Conn: c},
		Datacenters: &datacenters.Datacenters{Conn: c},
		Groups:      &groups.Groups{Conn: c},
		Users:       &users.Users{Conn: c},
	}
}
