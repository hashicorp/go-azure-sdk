// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2022_03_08_preview

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2022-03-08-preview/administrators"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2022-03-08-preview/backups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2022-03-08-preview/checknameavailability"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2022-03-08-preview/configurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2022-03-08-preview/databases"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2022-03-08-preview/firewallrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2022-03-08-preview/getprivatednszonesuffix"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2022-03-08-preview/locationbasedcapabilities"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2022-03-08-preview/replicas"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2022-03-08-preview/serverrestart"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2022-03-08-preview/servers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2022-03-08-preview/serverstart"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2022-03-08-preview/serverstop"
)

type Client struct {
	Administrators            *administrators.AdministratorsClient
	Backups                   *backups.BackupsClient
	CheckNameAvailability     *checknameavailability.CheckNameAvailabilityClient
	Configurations            *configurations.ConfigurationsClient
	Databases                 *databases.DatabasesClient
	FirewallRules             *firewallrules.FirewallRulesClient
	GetPrivateDnsZoneSuffix   *getprivatednszonesuffix.GetPrivateDnsZoneSuffixClient
	LocationBasedCapabilities *locationbasedcapabilities.LocationBasedCapabilitiesClient
	Replicas                  *replicas.ReplicasClient
	ServerRestart             *serverrestart.ServerRestartClient
	ServerStart               *serverstart.ServerStartClient
	ServerStop                *serverstop.ServerStopClient
	Servers                   *servers.ServersClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	administratorsClient := administrators.NewAdministratorsClientWithBaseURI(endpoint)
	configureAuthFunc(&administratorsClient.Client)

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

	replicasClient := replicas.NewReplicasClientWithBaseURI(endpoint)
	configureAuthFunc(&replicasClient.Client)

	serverRestartClient := serverrestart.NewServerRestartClientWithBaseURI(endpoint)
	configureAuthFunc(&serverRestartClient.Client)

	serverStartClient := serverstart.NewServerStartClientWithBaseURI(endpoint)
	configureAuthFunc(&serverStartClient.Client)

	serverStopClient := serverstop.NewServerStopClientWithBaseURI(endpoint)
	configureAuthFunc(&serverStopClient.Client)

	serversClient := servers.NewServersClientWithBaseURI(endpoint)
	configureAuthFunc(&serversClient.Client)

	return Client{
		Administrators:            &administratorsClient,
		Backups:                   &backupsClient,
		CheckNameAvailability:     &checkNameAvailabilityClient,
		Configurations:            &configurationsClient,
		Databases:                 &databasesClient,
		FirewallRules:             &firewallRulesClient,
		GetPrivateDnsZoneSuffix:   &getPrivateDnsZoneSuffixClient,
		LocationBasedCapabilities: &locationBasedCapabilitiesClient,
		Replicas:                  &replicasClient,
		ServerRestart:             &serverRestartClient,
		ServerStart:               &serverStartClient,
		ServerStop:                &serverStopClient,
		Servers:                   &serversClient,
	}
}
