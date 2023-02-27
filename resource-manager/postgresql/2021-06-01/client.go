// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2021_06_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2021-06-01/checknameavailability"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2021-06-01/configurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2021-06-01/databases"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2021-06-01/firewallrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2021-06-01/getprivatednszonesuffix"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2021-06-01/locationbasedcapabilities"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2021-06-01/serverrestart"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2021-06-01/servers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2021-06-01/serverstart"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2021-06-01/serverstop"
)

type Client struct {
	CheckNameAvailability     *checknameavailability.CheckNameAvailabilityClient
	Configurations            *configurations.ConfigurationsClient
	Databases                 *databases.DatabasesClient
	FirewallRules             *firewallrules.FirewallRulesClient
	GetPrivateDnsZoneSuffix   *getprivatednszonesuffix.GetPrivateDnsZoneSuffixClient
	LocationBasedCapabilities *locationbasedcapabilities.LocationBasedCapabilitiesClient
	ServerRestart             *serverrestart.ServerRestartClient
	ServerStart               *serverstart.ServerStartClient
	ServerStop                *serverstop.ServerStopClient
	Servers                   *servers.ServersClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	checkNameAvailabilityClient := checknameavailability.NewCheckNameAvailabilityClientWithBaseURI(endpoint)
	configureAuthFunc(&checkNameAvailabilityClient.Client)

	configurationsClient := configurations.NewConfigurationsClientWithBaseURI(endpoint)
	configureAuthFunc(&configurationsClient.Client)

	databasesClient := databases.NewDatabasesClientWithBaseURI(endpoint)
	configureAuthFunc(&databasesClient.Client)

	firewallRulesClient := firewallrules.NewFirewallRulesClientWithBaseURI(endpoint)
	configureAuthFunc(&firewallRulesClient.Client)

	getPrivateDnsZoneSuffixClient := getprivatednszonesuffix.NewGetPrivateDnsZoneSuffixClientWithBaseURI(endpoint)
	configureAuthFunc(&getPrivateDnsZoneSuffixClient.Client)

	locationBasedCapabilitiesClient := locationbasedcapabilities.NewLocationBasedCapabilitiesClientWithBaseURI(endpoint)
	configureAuthFunc(&locationBasedCapabilitiesClient.Client)

	serverRestartClient := serverrestart.NewServerRestartClientWithBaseURI(endpoint)
	configureAuthFunc(&serverRestartClient.Client)

	serverStartClient := serverstart.NewServerStartClientWithBaseURI(endpoint)
	configureAuthFunc(&serverStartClient.Client)

	serverStopClient := serverstop.NewServerStopClientWithBaseURI(endpoint)
	configureAuthFunc(&serverStopClient.Client)

	serversClient := servers.NewServersClientWithBaseURI(endpoint)
	configureAuthFunc(&serversClient.Client)

	return Client{
		CheckNameAvailability:     &checkNameAvailabilityClient,
		Configurations:            &configurationsClient,
		Databases:                 &databasesClient,
		FirewallRules:             &firewallRulesClient,
		GetPrivateDnsZoneSuffix:   &getPrivateDnsZoneSuffixClient,
		LocationBasedCapabilities: &locationBasedCapabilitiesClient,
		ServerRestart:             &serverRestartClient,
		ServerStart:               &serverStartClient,
		ServerStop:                &serverStopClient,
		Servers:                   &serversClient,
	}
}
