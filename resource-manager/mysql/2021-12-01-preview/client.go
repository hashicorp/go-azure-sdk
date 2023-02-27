// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2021_12_01_preview

import (
	"github.com/Azure/go-autorest/autorest"
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

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	azureADAdministratorsClient := azureadadministrators.NewAzureADAdministratorsClientWithBaseURI(endpoint)
	configureAuthFunc(&azureADAdministratorsClient.Client)

	backupsClient := backups.NewBackupsClientWithBaseURI(endpoint)
	configureAuthFunc(&backupsClient.Client)

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

	logFilesClient := logfiles.NewLogFilesClientWithBaseURI(endpoint)
	configureAuthFunc(&logFilesClient.Client)

	serverFailoverClient := serverfailover.NewServerFailoverClientWithBaseURI(endpoint)
	configureAuthFunc(&serverFailoverClient.Client)

	serverRestartClient := serverrestart.NewServerRestartClientWithBaseURI(endpoint)
	configureAuthFunc(&serverRestartClient.Client)

	serverStartClient := serverstart.NewServerStartClientWithBaseURI(endpoint)
	configureAuthFunc(&serverStartClient.Client)

	serverStopClient := serverstop.NewServerStopClientWithBaseURI(endpoint)
	configureAuthFunc(&serverStopClient.Client)

	serversClient := servers.NewServersClientWithBaseURI(endpoint)
	configureAuthFunc(&serversClient.Client)

	return Client{
		AzureADAdministrators:     &azureADAdministratorsClient,
		Backups:                   &backupsClient,
		CheckNameAvailability:     &checkNameAvailabilityClient,
		Configurations:            &configurationsClient,
		Databases:                 &databasesClient,
		FirewallRules:             &firewallRulesClient,
		GetPrivateDnsZoneSuffix:   &getPrivateDnsZoneSuffixClient,
		LocationBasedCapabilities: &locationBasedCapabilitiesClient,
		LogFiles:                  &logFilesClient,
		ServerFailover:            &serverFailoverClient,
		ServerRestart:             &serverRestartClient,
		ServerStart:               &serverStartClient,
		ServerStop:                &serverStopClient,
		Servers:                   &serversClient,
	}
}
