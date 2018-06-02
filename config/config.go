/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package config

import "net/url"

// Config : for storing credentials and information about ernest
type Config struct {
	Target           string
	Token            string
	Username         string
	Password         string
	VerificationCode string
	Version          string
}

// New : creates new config
func New(target string) *Config {
	return &Config{
		Target: target,
	}
}

// WithToken : sets the configs token
func (c *Config) WithToken(token string) *Config {
	c.Token = token
	return c
}

// WithCredentials : sets the configs credentials
func (c *Config) WithCredentials(username, password string) *Config {
	c.Username = username
	c.Password = password
	return c
}

// WithCredentialsAndVerification : sets the configs credentials
func (c *Config) WithCredentialsAndVerification(username, password, vc string) *Config {
	c.Username = username
	c.Password = password
	c.VerificationCode = vc
	return c
}

// Hostname : returns the hostname of the api
func (c *Config) Hostname() string {
	uri, _ := url.Parse(c.Target)
	if uri.Port() != "0" {
		return uri.Hostname() + ":" + uri.Port()
	}
	return uri.Hostname()
}

// WSScheme : returns the protocol (ws, wss) used for websocket connections
func (c *Config) WSScheme() string {
	uri, _ := url.Parse(c.Target)
	if uri.Scheme == "http" {
		return "ws"
	}
	return "wss"
}
