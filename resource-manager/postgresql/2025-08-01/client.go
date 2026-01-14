package v2025_08_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2025-08-01/administratormicrosoftentras"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2025-08-01/advancedthreatprotectionsettingsmodels"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2025-08-01/backupautomaticandondemands"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2025-08-01/backupslongtermretentionoperations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2025-08-01/configurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2025-08-01/databases"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2025-08-01/firewallrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2025-08-01/migrations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2025-08-01/openapis"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2025-08-01/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2025-08-01/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2025-08-01/servers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2025-08-01/tuningoptionsoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2025-08-01/virtualendpoints"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AdministratorMicrosoftEntras           *administratormicrosoftentras.AdministratorMicrosoftEntrasClient
	AdvancedThreatProtectionSettingsModels *advancedthreatprotectionsettingsmodels.AdvancedThreatProtectionSettingsModelsClient
	BackupAutomaticAndOnDemands            *backupautomaticandondemands.BackupAutomaticAndOnDemandsClient
	BackupsLongTermRetentionOperations     *backupslongtermretentionoperations.BackupsLongTermRetentionOperationsClient
	Configurations                         *configurations.ConfigurationsClient
	Databases                              *databases.DatabasesClient
	FirewallRules                          *firewallrules.FirewallRulesClient
	Migrations                             *migrations.MigrationsClient
	Openapis                               *openapis.OpenapisClient
	PrivateEndpointConnections             *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources                   *privatelinkresources.PrivateLinkResourcesClient
	Servers                                *servers.ServersClient
	TuningOptionsOperationGroup            *tuningoptionsoperationgroup.TuningOptionsOperationGroupClient
	VirtualEndpoints                       *virtualendpoints.VirtualEndpointsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	administratorMicrosoftEntrasClient, err := administratormicrosoftentras.NewAdministratorMicrosoftEntrasClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AdministratorMicrosoftEntras client: %+v", err)
	}
	configureFunc(administratorMicrosoftEntrasClient.Client)

	advancedThreatProtectionSettingsModelsClient, err := advancedthreatprotectionsettingsmodels.NewAdvancedThreatProtectionSettingsModelsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AdvancedThreatProtectionSettingsModels client: %+v", err)
	}
	configureFunc(advancedThreatProtectionSettingsModelsClient.Client)

	backupAutomaticAndOnDemandsClient, err := backupautomaticandondemands.NewBackupAutomaticAndOnDemandsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BackupAutomaticAndOnDemands client: %+v", err)
	}
	configureFunc(backupAutomaticAndOnDemandsClient.Client)

	backupsLongTermRetentionOperationsClient, err := backupslongtermretentionoperations.NewBackupsLongTermRetentionOperationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BackupsLongTermRetentionOperations client: %+v", err)
	}
	configureFunc(backupsLongTermRetentionOperationsClient.Client)

	configurationsClient, err := configurations.NewConfigurationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Configurations client: %+v", err)
	}
	configureFunc(configurationsClient.Client)

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

	migrationsClient, err := migrations.NewMigrationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Migrations client: %+v", err)
	}
	configureFunc(migrationsClient.Client)

	openapisClient, err := openapis.NewOpenapisClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Openapis client: %+v", err)
	}
	configureFunc(openapisClient.Client)

	privateEndpointConnectionsClient, err := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateEndpointConnections client: %+v", err)
	}
	configureFunc(privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient, err := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateLinkResources client: %+v", err)
	}
	configureFunc(privateLinkResourcesClient.Client)

	serversClient, err := servers.NewServersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Servers client: %+v", err)
	}
	configureFunc(serversClient.Client)

	tuningOptionsOperationGroupClient, err := tuningoptionsoperationgroup.NewTuningOptionsOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TuningOptionsOperationGroup client: %+v", err)
	}
	configureFunc(tuningOptionsOperationGroupClient.Client)

	virtualEndpointsClient, err := virtualendpoints.NewVirtualEndpointsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualEndpoints client: %+v", err)
	}
	configureFunc(virtualEndpointsClient.Client)

	return &Client{
		AdministratorMicrosoftEntras:           administratorMicrosoftEntrasClient,
		AdvancedThreatProtectionSettingsModels: advancedThreatProtectionSettingsModelsClient,
		BackupAutomaticAndOnDemands:            backupAutomaticAndOnDemandsClient,
		BackupsLongTermRetentionOperations:     backupsLongTermRetentionOperationsClient,
		Configurations:                         configurationsClient,
		Databases:                              databasesClient,
		FirewallRules:                          firewallRulesClient,
		Migrations:                             migrationsClient,
		Openapis:                               openapisClient,
		PrivateEndpointConnections:             privateEndpointConnectionsClient,
		PrivateLinkResources:                   privateLinkResourcesClient,
		Servers:                                serversClient,
		TuningOptionsOperationGroup:            tuningOptionsOperationGroupClient,
		VirtualEndpoints:                       virtualEndpointsClient,
	}, nil
}
