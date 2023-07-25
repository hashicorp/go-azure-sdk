package v2021_05_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/mysql/2021-05-01/backups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mysql/2021-05-01/checknameavailability"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mysql/2021-05-01/configurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mysql/2021-05-01/databases"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mysql/2021-05-01/firewallrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mysql/2021-05-01/getprivatednszonesuffix"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mysql/2021-05-01/locationbasedcapabilities"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mysql/2021-05-01/serverfailover"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mysql/2021-05-01/serverrestart"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mysql/2021-05-01/servers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mysql/2021-05-01/serverstart"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mysql/2021-05-01/serverstop"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Backups                   *backups.BackupsClient
	CheckNameAvailability     *checknameavailability.CheckNameAvailabilityClient
	Configurations            *configurations.ConfigurationsClient
	Databases                 *databases.DatabasesClient
	FirewallRules             *firewallrules.FirewallRulesClient
	GetPrivateDnsZoneSuffix   *getprivatednszonesuffix.GetPrivateDnsZoneSuffixClient
	LocationBasedCapabilities *locationbasedcapabilities.LocationBasedCapabilitiesClient
	ServerFailover            *serverfailover.ServerFailoverClient
	ServerRestart             *serverrestart.ServerRestartClient
	ServerStart               *serverstart.ServerStartClient
	ServerStop                *serverstop.ServerStopClient
	Servers                   *servers.ServersClient
}

func NewClientWithBaseURI(api environments.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	backupsClient, err := backups.NewBackupsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Backups client: %+v", err)
	}
	configureFunc(backupsClient.Client)

	checkNameAvailabilityClient, err := checknameavailability.NewCheckNameAvailabilityClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building CheckNameAvailability client: %+v", err)
	}
	configureFunc(checkNameAvailabilityClient.Client)

	configurationsClient, err := configurations.NewConfigurationsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Configurations client: %+v", err)
	}
	configureFunc(configurationsClient.Client)

	databasesClient, err := databases.NewDatabasesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Databases client: %+v", err)
	}
	configureFunc(databasesClient.Client)

	firewallRulesClient, err := firewallrules.NewFirewallRulesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building FirewallRules client: %+v", err)
	}
	configureFunc(firewallRulesClient.Client)

	getPrivateDnsZoneSuffixClient, err := getprivatednszonesuffix.NewGetPrivateDnsZoneSuffixClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building GetPrivateDnsZoneSuffix client: %+v", err)
	}
	configureFunc(getPrivateDnsZoneSuffixClient.Client)

	locationBasedCapabilitiesClient, err := locationbasedcapabilities.NewLocationBasedCapabilitiesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building LocationBasedCapabilities client: %+v", err)
	}
	configureFunc(locationBasedCapabilitiesClient.Client)

	serverFailoverClient, err := serverfailover.NewServerFailoverClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ServerFailover client: %+v", err)
	}
	configureFunc(serverFailoverClient.Client)

	serverRestartClient, err := serverrestart.NewServerRestartClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ServerRestart client: %+v", err)
	}
	configureFunc(serverRestartClient.Client)

	serverStartClient, err := serverstart.NewServerStartClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ServerStart client: %+v", err)
	}
	configureFunc(serverStartClient.Client)

	serverStopClient, err := serverstop.NewServerStopClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ServerStop client: %+v", err)
	}
	configureFunc(serverStopClient.Client)

	serversClient, err := servers.NewServersClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Servers client: %+v", err)
	}
	configureFunc(serversClient.Client)

	return &Client{
		Backups:                   backupsClient,
		CheckNameAvailability:     checkNameAvailabilityClient,
		Configurations:            configurationsClient,
		Databases:                 databasesClient,
		FirewallRules:             firewallRulesClient,
		GetPrivateDnsZoneSuffix:   getPrivateDnsZoneSuffixClient,
		LocationBasedCapabilities: locationBasedCapabilitiesClient,
		ServerFailover:            serverFailoverClient,
		ServerRestart:             serverRestartClient,
		ServerStart:               serverStartClient,
		ServerStop:                serverStopClient,
		Servers:                   serversClient,
	}, nil
}
