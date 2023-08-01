package v2016_11_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/datalakeanalytics/2016-11-01/accounts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datalakeanalytics/2016-11-01/computepolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datalakeanalytics/2016-11-01/datalakestoreaccounts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datalakeanalytics/2016-11-01/firewallrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datalakeanalytics/2016-11-01/locations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datalakeanalytics/2016-11-01/storageaccounts"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Accounts              *accounts.AccountsClient
	ComputePolicies       *computepolicies.ComputePoliciesClient
	DataLakeStoreAccounts *datalakestoreaccounts.DataLakeStoreAccountsClient
	FirewallRules         *firewallrules.FirewallRulesClient
	Locations             *locations.LocationsClient
	StorageAccounts       *storageaccounts.StorageAccountsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	accountsClient, err := accounts.NewAccountsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Accounts client: %+v", err)
	}
	configureFunc(accountsClient.Client)

	computePoliciesClient, err := computepolicies.NewComputePoliciesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ComputePolicies client: %+v", err)
	}
	configureFunc(computePoliciesClient.Client)

	dataLakeStoreAccountsClient, err := datalakestoreaccounts.NewDataLakeStoreAccountsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DataLakeStoreAccounts client: %+v", err)
	}
	configureFunc(dataLakeStoreAccountsClient.Client)

	firewallRulesClient, err := firewallrules.NewFirewallRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FirewallRules client: %+v", err)
	}
	configureFunc(firewallRulesClient.Client)

	locationsClient, err := locations.NewLocationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Locations client: %+v", err)
	}
	configureFunc(locationsClient.Client)

	storageAccountsClient, err := storageaccounts.NewStorageAccountsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building StorageAccounts client: %+v", err)
	}
	configureFunc(storageAccountsClient.Client)

	return &Client{
		Accounts:              accountsClient,
		ComputePolicies:       computePoliciesClient,
		DataLakeStoreAccounts: dataLakeStoreAccountsClient,
		FirewallRules:         firewallRulesClient,
		Locations:             locationsClient,
		StorageAccounts:       storageAccountsClient,
	}, nil
}
