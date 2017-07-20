/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

// Datacenter : stores datacenter data
type Datacenter struct {
	ID              int    `json:"id"`
	GroupID         int    `json:"group_id"`
	Name            string `json:"name"`
	Type            string `json:"type"`
	Region          string `json:"region,omitempty"`
	Username        string `json:"username,omitempty"`
	Password        string `json:"password,omitempty"`
	VCloudURL       string `json:"vcloud_url,omitempty"`
	VseURL          string `json:"vse_url,omitempty"`
	ExternalNetwork string `json:"external_network,omitempty"`
	AccessKeyID     string `json:"aws_access_key_id,omitempty"`
	SecretAccessKey string `json:"aws_secret_access_key,omitempty"`
	SubscriptionID  string `json:"azure_subscription_id,omitempty"`
	ClientID        string `json:"azure_client_id,omitempty"`
	ClientSecret    string `json:"azure_client_secret,omitempty"`
	TenantID        string `json:"azure_tenant_id"`
	Environment     string `json:"azure_environment"`
}
