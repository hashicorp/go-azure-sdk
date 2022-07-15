package v2016_11_01

import (
	"github.com/hashicorp/go-azure-sdk/resource-manager/datalakestore/2016-11-01/accounts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datalakestore/2016-11-01/firewallrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datalakestore/2016-11-01/locations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datalakestore/2016-11-01/trustedidproviders"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datalakestore/2016-11-01/virtualnetworkrules"
)

type Client struct {
	Accounts            *accounts.AccountsClient
	FirewallRules       *firewallrules.FirewallRulesClient
	Locations           *locations.LocationsClient
	TrustedIdProviders  *trustedidproviders.TrustedIdProvidersClient
	VirtualNetworkRules *virtualnetworkrules.VirtualNetworkRulesClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	accountsClient := accounts.NewAccountsClientWithBaseURI(endpoint)
	configureAuthFunc(&accountsClient.Client)

	firewallRulesClient := firewallrules.NewFirewallRulesClientWithBaseURI(endpoint)
	configureAuthFunc(&firewallRulesClient.Client)

	locationsClient := locations.NewLocationsClientWithBaseURI(endpoint)
	configureAuthFunc(&locationsClient.Client)

	trustedIdProvidersClient := trustedidproviders.NewTrustedIdProvidersClientWithBaseURI(endpoint)
	configureAuthFunc(&trustedIdProvidersClient.Client)

	virtualNetworkRulesClient := virtualnetworkrules.NewVirtualNetworkRulesClientWithBaseURI(endpoint)
	configureAuthFunc(&virtualNetworkRulesClient.Client)

	return Client{
		Accounts:            &accountsClient,
		FirewallRules:       &firewallRulesClient,
		Locations:           &locationsClient,
		TrustedIdProviders:  &trustedIdProvidersClient,
		VirtualNetworkRules: &virtualNetworkRulesClient,
	}
}
