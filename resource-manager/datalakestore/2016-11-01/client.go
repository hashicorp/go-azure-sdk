package v2016_11_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/datalakestore/2016-11-01/accounts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datalakestore/2016-11-01/firewallrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datalakestore/2016-11-01/locations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datalakestore/2016-11-01/trustedidproviders"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datalakestore/2016-11-01/virtualnetworkrules"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Accounts            *accounts.AccountsClient
	FirewallRules       *firewallrules.FirewallRulesClient
	Locations           *locations.LocationsClient
	TrustedIdProviders  *trustedidproviders.TrustedIdProvidersClient
	VirtualNetworkRules *virtualnetworkrules.VirtualNetworkRulesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	accountsClient, err := accounts.NewAccountsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Accounts client: %+v", err)
	}
	configureFunc(accountsClient.Client)

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

	trustedIdProvidersClient, err := trustedidproviders.NewTrustedIdProvidersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TrustedIdProviders client: %+v", err)
	}
	configureFunc(trustedIdProvidersClient.Client)

	virtualNetworkRulesClient, err := virtualnetworkrules.NewVirtualNetworkRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualNetworkRules client: %+v", err)
	}
	configureFunc(virtualNetworkRulesClient.Client)

	return &Client{
		Accounts:            accountsClient,
		FirewallRules:       firewallRulesClient,
		Locations:           locationsClient,
		TrustedIdProviders:  trustedIdProvidersClient,
		VirtualNetworkRules: virtualNetworkRulesClient,
	}, nil
}
