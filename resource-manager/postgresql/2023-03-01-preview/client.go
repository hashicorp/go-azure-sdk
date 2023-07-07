package v2023_03_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2023-03-01-preview/administrators"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2023-03-01-preview/backups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2023-03-01-preview/checknameavailability"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2023-03-01-preview/configurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2023-03-01-preview/customoperation"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2023-03-01-preview/databases"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2023-03-01-preview/firewallrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2023-03-01-preview/flexibleservercapabilities"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2023-03-01-preview/getprivatednszonesuffix"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2023-03-01-preview/locationbasedcapabilities"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2023-03-01-preview/logfiles"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2023-03-01-preview/longtermretentionbackup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2023-03-01-preview/migrations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2023-03-01-preview/post"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2023-03-01-preview/replicas"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2023-03-01-preview/serverrestart"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2023-03-01-preview/servers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2023-03-01-preview/serverstart"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2023-03-01-preview/serverstop"
)

type Client struct {
	Administrators             *administrators.AdministratorsClient
	Backups                    *backups.BackupsClient
	CheckNameAvailability      *checknameavailability.CheckNameAvailabilityClient
	Configurations             *configurations.ConfigurationsClient
	CustomOperation            *customoperation.CustomOperationClient
	Databases                  *databases.DatabasesClient
	FirewallRules              *firewallrules.FirewallRulesClient
	FlexibleServerCapabilities *flexibleservercapabilities.FlexibleServerCapabilitiesClient
	GetPrivateDnsZoneSuffix    *getprivatednszonesuffix.GetPrivateDnsZoneSuffixClient
	LocationBasedCapabilities  *locationbasedcapabilities.LocationBasedCapabilitiesClient
	LogFiles                   *logfiles.LogFilesClient
	LongTermRetentionBackup    *longtermretentionbackup.LongTermRetentionBackupClient
	Migrations                 *migrations.MigrationsClient
	POST                       *post.POSTClient
	Replicas                   *replicas.ReplicasClient
	ServerRestart              *serverrestart.ServerRestartClient
	ServerStart                *serverstart.ServerStartClient
	ServerStop                 *serverstop.ServerStopClient
	Servers                    *servers.ServersClient
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

	customOperationClient := customoperation.NewCustomOperationClientWithBaseURI(endpoint)
	configureAuthFunc(&customOperationClient.Client)

	databasesClient := databases.NewDatabasesClientWithBaseURI(endpoint)
	configureAuthFunc(&databasesClient.Client)

	firewallRulesClient := firewallrules.NewFirewallRulesClientWithBaseURI(endpoint)
	configureAuthFunc(&firewallRulesClient.Client)

	flexibleServerCapabilitiesClient := flexibleservercapabilities.NewFlexibleServerCapabilitiesClientWithBaseURI(endpoint)
	configureAuthFunc(&flexibleServerCapabilitiesClient.Client)

	getPrivateDnsZoneSuffixClient := getprivatednszonesuffix.NewGetPrivateDnsZoneSuffixClientWithBaseURI(endpoint)
	configureAuthFunc(&getPrivateDnsZoneSuffixClient.Client)

	locationBasedCapabilitiesClient := locationbasedcapabilities.NewLocationBasedCapabilitiesClientWithBaseURI(endpoint)
	configureAuthFunc(&locationBasedCapabilitiesClient.Client)

	logFilesClient := logfiles.NewLogFilesClientWithBaseURI(endpoint)
	configureAuthFunc(&logFilesClient.Client)

	longTermRetentionBackupClient := longtermretentionbackup.NewLongTermRetentionBackupClientWithBaseURI(endpoint)
	configureAuthFunc(&longTermRetentionBackupClient.Client)

	migrationsClient := migrations.NewMigrationsClientWithBaseURI(endpoint)
	configureAuthFunc(&migrationsClient.Client)

	pOSTClient := post.NewPOSTClientWithBaseURI(endpoint)
	configureAuthFunc(&pOSTClient.Client)

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
		Administrators:             &administratorsClient,
		Backups:                    &backupsClient,
		CheckNameAvailability:      &checkNameAvailabilityClient,
		Configurations:             &configurationsClient,
		CustomOperation:            &customOperationClient,
		Databases:                  &databasesClient,
		FirewallRules:              &firewallRulesClient,
		FlexibleServerCapabilities: &flexibleServerCapabilitiesClient,
		GetPrivateDnsZoneSuffix:    &getPrivateDnsZoneSuffixClient,
		LocationBasedCapabilities:  &locationBasedCapabilitiesClient,
		LogFiles:                   &logFilesClient,
		LongTermRetentionBackup:    &longTermRetentionBackupClient,
		Migrations:                 &migrationsClient,
		POST:                       &pOSTClient,
		Replicas:                   &replicasClient,
		ServerRestart:              &serverRestartClient,
		ServerStart:                &serverStartClient,
		ServerStop:                 &serverStopClient,
		Servers:                    &serversClient,
	}
}
