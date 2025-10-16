package v2025_08_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2025-08-01/administratorsmicrosoftentra"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2025-08-01/advancedthreatprotectionsettings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2025-08-01/backupsautomaticandondemand"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2025-08-01/backupslongtermretention"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2025-08-01/capabilitiesbylocation"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2025-08-01/capabilitiesbyserver"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2025-08-01/capturedlogs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2025-08-01/configurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2025-08-01/customoperation"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2025-08-01/databases"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2025-08-01/firewallrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2025-08-01/migrations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2025-08-01/nameavailability"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2025-08-01/post"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2025-08-01/privatednszonesuffix"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2025-08-01/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2025-08-01/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2025-08-01/quotausagesforflexibleservers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2025-08-01/replicas"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2025-08-01/servers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2025-08-01/tuningoptions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2025-08-01/virtualendpoints"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AdministratorsMicrosoftEntra     *administratorsmicrosoftentra.AdministratorsMicrosoftEntraClient
	AdvancedThreatProtectionSettings *advancedthreatprotectionsettings.AdvancedThreatProtectionSettingsClient
	BackupsAutomaticAndOnDemand      *backupsautomaticandondemand.BackupsAutomaticAndOnDemandClient
	BackupsLongTermRetention         *backupslongtermretention.BackupsLongTermRetentionClient
	CapabilitiesByLocation           *capabilitiesbylocation.CapabilitiesByLocationClient
	CapabilitiesByServer             *capabilitiesbyserver.CapabilitiesByServerClient
	CapturedLogs                     *capturedlogs.CapturedLogsClient
	Configurations                   *configurations.ConfigurationsClient
	CustomOperation                  *customoperation.CustomOperationClient
	Databases                        *databases.DatabasesClient
	FirewallRules                    *firewallrules.FirewallRulesClient
	Migrations                       *migrations.MigrationsClient
	NameAvailability                 *nameavailability.NameAvailabilityClient
	POST                             *post.POSTClient
	PrivateDnsZoneSuffix             *privatednszonesuffix.PrivateDnsZoneSuffixClient
	PrivateEndpointConnections       *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources             *privatelinkresources.PrivateLinkResourcesClient
	QuotaUsagesForFlexibleServers    *quotausagesforflexibleservers.QuotaUsagesForFlexibleServersClient
	Replicas                         *replicas.ReplicasClient
	Servers                          *servers.ServersClient
	TuningOptions                    *tuningoptions.TuningOptionsClient
	VirtualEndpoints                 *virtualendpoints.VirtualEndpointsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	administratorsMicrosoftEntraClient, err := administratorsmicrosoftentra.NewAdministratorsMicrosoftEntraClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AdministratorsMicrosoftEntra client: %+v", err)
	}
	configureFunc(administratorsMicrosoftEntraClient.Client)

	advancedThreatProtectionSettingsClient, err := advancedthreatprotectionsettings.NewAdvancedThreatProtectionSettingsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AdvancedThreatProtectionSettings client: %+v", err)
	}
	configureFunc(advancedThreatProtectionSettingsClient.Client)

	backupsAutomaticAndOnDemandClient, err := backupsautomaticandondemand.NewBackupsAutomaticAndOnDemandClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BackupsAutomaticAndOnDemand client: %+v", err)
	}
	configureFunc(backupsAutomaticAndOnDemandClient.Client)

	backupsLongTermRetentionClient, err := backupslongtermretention.NewBackupsLongTermRetentionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BackupsLongTermRetention client: %+v", err)
	}
	configureFunc(backupsLongTermRetentionClient.Client)

	capabilitiesByLocationClient, err := capabilitiesbylocation.NewCapabilitiesByLocationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CapabilitiesByLocation client: %+v", err)
	}
	configureFunc(capabilitiesByLocationClient.Client)

	capabilitiesByServerClient, err := capabilitiesbyserver.NewCapabilitiesByServerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CapabilitiesByServer client: %+v", err)
	}
	configureFunc(capabilitiesByServerClient.Client)

	capturedLogsClient, err := capturedlogs.NewCapturedLogsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CapturedLogs client: %+v", err)
	}
	configureFunc(capturedLogsClient.Client)

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

	migrationsClient, err := migrations.NewMigrationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Migrations client: %+v", err)
	}
	configureFunc(migrationsClient.Client)

	nameAvailabilityClient, err := nameavailability.NewNameAvailabilityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NameAvailability client: %+v", err)
	}
	configureFunc(nameAvailabilityClient.Client)

	pOSTClient, err := post.NewPOSTClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building POST client: %+v", err)
	}
	configureFunc(pOSTClient.Client)

	privateDnsZoneSuffixClient, err := privatednszonesuffix.NewPrivateDnsZoneSuffixClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateDnsZoneSuffix client: %+v", err)
	}
	configureFunc(privateDnsZoneSuffixClient.Client)

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

	quotaUsagesForFlexibleServersClient, err := quotausagesforflexibleservers.NewQuotaUsagesForFlexibleServersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building QuotaUsagesForFlexibleServers client: %+v", err)
	}
	configureFunc(quotaUsagesForFlexibleServersClient.Client)

	replicasClient, err := replicas.NewReplicasClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Replicas client: %+v", err)
	}
	configureFunc(replicasClient.Client)

	serversClient, err := servers.NewServersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Servers client: %+v", err)
	}
	configureFunc(serversClient.Client)

	tuningOptionsClient, err := tuningoptions.NewTuningOptionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TuningOptions client: %+v", err)
	}
	configureFunc(tuningOptionsClient.Client)

	virtualEndpointsClient, err := virtualendpoints.NewVirtualEndpointsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualEndpoints client: %+v", err)
	}
	configureFunc(virtualEndpointsClient.Client)

	return &Client{
		AdministratorsMicrosoftEntra:     administratorsMicrosoftEntraClient,
		AdvancedThreatProtectionSettings: advancedThreatProtectionSettingsClient,
		BackupsAutomaticAndOnDemand:      backupsAutomaticAndOnDemandClient,
		BackupsLongTermRetention:         backupsLongTermRetentionClient,
		CapabilitiesByLocation:           capabilitiesByLocationClient,
		CapabilitiesByServer:             capabilitiesByServerClient,
		CapturedLogs:                     capturedLogsClient,
		Configurations:                   configurationsClient,
		CustomOperation:                  customOperationClient,
		Databases:                        databasesClient,
		FirewallRules:                    firewallRulesClient,
		Migrations:                       migrationsClient,
		NameAvailability:                 nameAvailabilityClient,
		POST:                             pOSTClient,
		PrivateDnsZoneSuffix:             privateDnsZoneSuffixClient,
		PrivateEndpointConnections:       privateEndpointConnectionsClient,
		PrivateLinkResources:             privateLinkResourcesClient,
		QuotaUsagesForFlexibleServers:    quotaUsagesForFlexibleServersClient,
		Replicas:                         replicasClient,
		Servers:                          serversClient,
		TuningOptions:                    tuningOptionsClient,
		VirtualEndpoints:                 virtualEndpointsClient,
	}, nil
}
