package v2023_03_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

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
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
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

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	administratorsClient, err := administrators.NewAdministratorsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Administrators client: %+v", err)
	}
	configureFunc(administratorsClient.Client)

	backupsClient, err := backups.NewBackupsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Backups client: %+v", err)
	}
	configureFunc(backupsClient.Client)

	checkNameAvailabilityClient, err := checknameavailability.NewCheckNameAvailabilityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CheckNameAvailability client: %+v", err)
	}
	configureFunc(checkNameAvailabilityClient.Client)

	configurationsClient, err := configurations.NewConfigurationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Configurations client: %+v", err)
	}
	configureFunc(configurationsClient.Client)

	customOperationClient, err := customoperation.NewCustomOperationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CustomOperation client: %+v", err)
	}
	configureFunc(customOperationClient.Client)

	databasesClient, err := databases.NewDatabasesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Databases client: %+v", err)
	}
	configureFunc(databasesClient.Client)

	firewallRulesClient, err := firewallrules.NewFirewallRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FirewallRules client: %+v", err)
	}
	configureFunc(firewallRulesClient.Client)

	flexibleServerCapabilitiesClient, err := flexibleservercapabilities.NewFlexibleServerCapabilitiesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FlexibleServerCapabilities client: %+v", err)
	}
	configureFunc(flexibleServerCapabilitiesClient.Client)

	getPrivateDnsZoneSuffixClient, err := getprivatednszonesuffix.NewGetPrivateDnsZoneSuffixClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GetPrivateDnsZoneSuffix client: %+v", err)
	}
	configureFunc(getPrivateDnsZoneSuffixClient.Client)

	locationBasedCapabilitiesClient, err := locationbasedcapabilities.NewLocationBasedCapabilitiesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LocationBasedCapabilities client: %+v", err)
	}
	configureFunc(locationBasedCapabilitiesClient.Client)

	logFilesClient, err := logfiles.NewLogFilesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LogFiles client: %+v", err)
	}
	configureFunc(logFilesClient.Client)

	longTermRetentionBackupClient, err := longtermretentionbackup.NewLongTermRetentionBackupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LongTermRetentionBackup client: %+v", err)
	}
	configureFunc(longTermRetentionBackupClient.Client)

	migrationsClient, err := migrations.NewMigrationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Migrations client: %+v", err)
	}
	configureFunc(migrationsClient.Client)

	pOSTClient, err := post.NewPOSTClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building POST client: %+v", err)
	}
	configureFunc(pOSTClient.Client)

	replicasClient, err := replicas.NewReplicasClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Replicas client: %+v", err)
	}
	configureFunc(replicasClient.Client)

	serverRestartClient, err := serverrestart.NewServerRestartClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServerRestart client: %+v", err)
	}
	configureFunc(serverRestartClient.Client)

	serverStartClient, err := serverstart.NewServerStartClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServerStart client: %+v", err)
	}
	configureFunc(serverStartClient.Client)

	serverStopClient, err := serverstop.NewServerStopClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServerStop client: %+v", err)
	}
	configureFunc(serverStopClient.Client)

	serversClient, err := servers.NewServersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Servers client: %+v", err)
	}
	configureFunc(serversClient.Client)

	return &Client{
		Administrators:             administratorsClient,
		Backups:                    backupsClient,
		CheckNameAvailability:      checkNameAvailabilityClient,
		Configurations:             configurationsClient,
		CustomOperation:            customOperationClient,
		Databases:                  databasesClient,
		FirewallRules:              firewallRulesClient,
		FlexibleServerCapabilities: flexibleServerCapabilitiesClient,
		GetPrivateDnsZoneSuffix:    getPrivateDnsZoneSuffixClient,
		LocationBasedCapabilities:  locationBasedCapabilitiesClient,
		LogFiles:                   logFilesClient,
		LongTermRetentionBackup:    longTermRetentionBackupClient,
		Migrations:                 migrationsClient,
		POST:                       pOSTClient,
		Replicas:                   replicasClient,
		ServerRestart:              serverRestartClient,
		ServerStart:                serverStartClient,
		ServerStop:                 serverStopClient,
		Servers:                    serversClient,
	}, nil
}
