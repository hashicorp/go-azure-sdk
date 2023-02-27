// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2016_11_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datalakeanalytics/2016-11-01/accounts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datalakeanalytics/2016-11-01/computepolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datalakeanalytics/2016-11-01/datalakestoreaccounts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datalakeanalytics/2016-11-01/firewallrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datalakeanalytics/2016-11-01/locations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datalakeanalytics/2016-11-01/storageaccounts"
)

type Client struct {
	Accounts              *accounts.AccountsClient
	ComputePolicies       *computepolicies.ComputePoliciesClient
	DataLakeStoreAccounts *datalakestoreaccounts.DataLakeStoreAccountsClient
	FirewallRules         *firewallrules.FirewallRulesClient
	Locations             *locations.LocationsClient
	StorageAccounts       *storageaccounts.StorageAccountsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	accountsClient := accounts.NewAccountsClientWithBaseURI(endpoint)
	configureAuthFunc(&accountsClient.Client)

	computePoliciesClient := computepolicies.NewComputePoliciesClientWithBaseURI(endpoint)
	configureAuthFunc(&computePoliciesClient.Client)

	dataLakeStoreAccountsClient := datalakestoreaccounts.NewDataLakeStoreAccountsClientWithBaseURI(endpoint)
	configureAuthFunc(&dataLakeStoreAccountsClient.Client)

	firewallRulesClient := firewallrules.NewFirewallRulesClientWithBaseURI(endpoint)
	configureAuthFunc(&firewallRulesClient.Client)

	locationsClient := locations.NewLocationsClientWithBaseURI(endpoint)
	configureAuthFunc(&locationsClient.Client)

	storageAccountsClient := storageaccounts.NewStorageAccountsClientWithBaseURI(endpoint)
	configureAuthFunc(&storageAccountsClient.Client)

	return Client{
		Accounts:              &accountsClient,
		ComputePolicies:       &computePoliciesClient,
		DataLakeStoreAccounts: &dataLakeStoreAccountsClient,
		FirewallRules:         &firewallRulesClient,
		Locations:             &locationsClient,
		StorageAccounts:       &storageAccountsClient,
	}
}
