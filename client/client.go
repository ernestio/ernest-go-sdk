/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package client

import (
	"github.com/ernestio/ernest-sdk/client/services"
	"github.com/ernestio/ernest-sdk/config"
	"github.com/ernestio/ernest-sdk/connection"
)

// Client :
type Client struct {
	Conn     *connection.Conn
	Services *services.Services
}

// New : creates a new client
func New(cfg *config.Config) *Client {
	c := connection.New(cfg)

	return &Client{
		Conn:     c,
		Services: &services.Services{Conn: c},
	}
}
