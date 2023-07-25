package v2021_12_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/mysql/2021-12-01-preview/azureadadministrators"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mysql/2021-12-01-preview/backups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mysql/2021-12-01-preview/checknameavailability"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mysql/2021-12-01-preview/configurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mysql/2021-12-01-preview/databases"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mysql/2021-12-01-preview/firewallrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mysql/2021-12-01-preview/getprivatednszonesuffix"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mysql/2021-12-01-preview/locationbasedcapabilities"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mysql/2021-12-01-preview/logfiles"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mysql/2021-12-01-preview/serverfailover"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mysql/2021-12-01-preview/serverrestart"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mysql/2021-12-01-preview/servers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mysql/2021-12-01-preview/serverstart"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mysql/2021-12-01-preview/serverstop"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AzureADAdministrators     *azureadadministrators.AzureADAdministratorsClient
	Backups                   *backups.BackupsClient
	CheckNameAvailability     *checknameavailability.CheckNameAvailabilityClient
	Configurations            *configurations.ConfigurationsClient
	Databases                 *databases.DatabasesClient
	FirewallRules             *firewallrules.FirewallRulesClient
	GetPrivateDnsZoneSuffix   *getprivatednszonesuffix.GetPrivateDnsZoneSuffixClient
	LocationBasedCapabilities *locationbasedcapabilities.LocationBasedCapabilitiesClient
	LogFiles                  *logfiles.LogFilesClient
	ServerFailover            *serverfailover.ServerFailoverClient
	ServerRestart             *serverrestart.ServerRestartClient
	ServerStart               *serverstart.ServerStartClient
	ServerStop                *serverstop.ServerStopClient
	Servers                   *servers.ServersClient
}

func NewClientWithBaseURI(api environments.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	azureADAdministratorsClient, err := azureadadministrators.NewAzureADAdministratorsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building AzureADAdministrators client: %+v", err)
	}
	configureFunc(azureADAdministratorsClient.Client)

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

	logFilesClient, err := logfiles.NewLogFilesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building LogFiles client: %+v", err)
	}
	configureFunc(logFilesClient.Client)

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
		AzureADAdministrators:     azureADAdministratorsClient,
		Backups:                   backupsClient,
		CheckNameAvailability:     checkNameAvailabilityClient,
		Configurations:            configurationsClient,
		Databases:                 databasesClient,
		FirewallRules:             firewallRulesClient,
		GetPrivateDnsZoneSuffix:   getPrivateDnsZoneSuffixClient,
		LocationBasedCapabilities: locationBasedCapabilitiesClient,
		LogFiles:                  logFilesClient,
		ServerFailover:            serverFailoverClient,
		ServerRestart:             serverRestartClient,
		ServerStart:               serverStartClient,
		ServerStop:                serverStopClient,
		Servers:                   serversClient,
	}, nil
}
