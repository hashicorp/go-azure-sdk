package v2026_10_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/horizondb/2026-10-01-preview/authentications"
	"github.com/hashicorp/go-azure-sdk/resource-manager/horizondb/2026-10-01-preview/capabilities"
	"github.com/hashicorp/go-azure-sdk/resource-manager/horizondb/2026-10-01-preview/changedatacapturecapabilities"
	"github.com/hashicorp/go-azure-sdk/resource-manager/horizondb/2026-10-01-preview/changedatacaptureconfigurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/horizondb/2026-10-01-preview/changedatacapturedestinations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/horizondb/2026-10-01-preview/clusters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/horizondb/2026-10-01-preview/defaultparametergroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/horizondb/2026-10-01-preview/endpoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/horizondb/2026-10-01-preview/firewallrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/horizondb/2026-10-01-preview/identityassignments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/horizondb/2026-10-01-preview/logcaptureconfigurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/horizondb/2026-10-01-preview/logfiles"
	"github.com/hashicorp/go-azure-sdk/resource-manager/horizondb/2026-10-01-preview/maintenanceevents"
	"github.com/hashicorp/go-azure-sdk/resource-manager/horizondb/2026-10-01-preview/microsoftentraadministrators"
	"github.com/hashicorp/go-azure-sdk/resource-manager/horizondb/2026-10-01-preview/microsoftentraauthentications"
	"github.com/hashicorp/go-azure-sdk/resource-manager/horizondb/2026-10-01-preview/nodes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/horizondb/2026-10-01-preview/parametergroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/horizondb/2026-10-01-preview/passwordauthentications"
	"github.com/hashicorp/go-azure-sdk/resource-manager/horizondb/2026-10-01-preview/pools"
	"github.com/hashicorp/go-azure-sdk/resource-manager/horizondb/2026-10-01-preview/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/horizondb/2026-10-01-preview/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Authentications                 *authentications.AuthenticationsClient
	Capabilities                    *capabilities.CapabilitiesClient
	ChangeDataCaptureCapabilities   *changedatacapturecapabilities.ChangeDataCaptureCapabilitiesClient
	ChangeDataCaptureConfigurations *changedatacaptureconfigurations.ChangeDataCaptureConfigurationsClient
	ChangeDataCaptureDestinations   *changedatacapturedestinations.ChangeDataCaptureDestinationsClient
	Clusters                        *clusters.ClustersClient
	DefaultParameterGroups          *defaultparametergroups.DefaultParameterGroupsClient
	Endpoints                       *endpoints.EndpointsClient
	FirewallRules                   *firewallrules.FirewallRulesClient
	IdentityAssignments             *identityassignments.IdentityAssignmentsClient
	LogCaptureConfigurations        *logcaptureconfigurations.LogCaptureConfigurationsClient
	LogFiles                        *logfiles.LogFilesClient
	MaintenanceEvents               *maintenanceevents.MaintenanceEventsClient
	MicrosoftEntraAdministrators    *microsoftentraadministrators.MicrosoftEntraAdministratorsClient
	MicrosoftEntraAuthentications   *microsoftentraauthentications.MicrosoftEntraAuthenticationsClient
	Nodes                           *nodes.NodesClient
	ParameterGroups                 *parametergroups.ParameterGroupsClient
	PasswordAuthentications         *passwordauthentications.PasswordAuthenticationsClient
	Pools                           *pools.PoolsClient
	PrivateEndpointConnections      *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources            *privatelinkresources.PrivateLinkResourcesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	authenticationsClient, err := authentications.NewAuthenticationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Authentications client: %+v", err)
	}
	configureFunc(authenticationsClient.Client)

	capabilitiesClient, err := capabilities.NewCapabilitiesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Capabilities client: %+v", err)
	}
	configureFunc(capabilitiesClient.Client)

	changeDataCaptureCapabilitiesClient, err := changedatacapturecapabilities.NewChangeDataCaptureCapabilitiesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ChangeDataCaptureCapabilities client: %+v", err)
	}
	configureFunc(changeDataCaptureCapabilitiesClient.Client)

	changeDataCaptureConfigurationsClient, err := changedatacaptureconfigurations.NewChangeDataCaptureConfigurationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ChangeDataCaptureConfigurations client: %+v", err)
	}
	configureFunc(changeDataCaptureConfigurationsClient.Client)

	changeDataCaptureDestinationsClient, err := changedatacapturedestinations.NewChangeDataCaptureDestinationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ChangeDataCaptureDestinations client: %+v", err)
	}
	configureFunc(changeDataCaptureDestinationsClient.Client)

	clustersClient, err := clusters.NewClustersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Clusters client: %+v", err)
	}
	configureFunc(clustersClient.Client)

	defaultParameterGroupsClient, err := defaultparametergroups.NewDefaultParameterGroupsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DefaultParameterGroups client: %+v", err)
	}
	configureFunc(defaultParameterGroupsClient.Client)

	endpointsClient, err := endpoints.NewEndpointsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Endpoints client: %+v", err)
	}
	configureFunc(endpointsClient.Client)

	firewallRulesClient, err := firewallrules.NewFirewallRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FirewallRules client: %+v", err)
	}
	configureFunc(firewallRulesClient.Client)

	identityAssignmentsClient, err := identityassignments.NewIdentityAssignmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IdentityAssignments client: %+v", err)
	}
	configureFunc(identityAssignmentsClient.Client)

	logCaptureConfigurationsClient, err := logcaptureconfigurations.NewLogCaptureConfigurationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LogCaptureConfigurations client: %+v", err)
	}
	configureFunc(logCaptureConfigurationsClient.Client)

	logFilesClient, err := logfiles.NewLogFilesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LogFiles client: %+v", err)
	}
	configureFunc(logFilesClient.Client)

	maintenanceEventsClient, err := maintenanceevents.NewMaintenanceEventsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MaintenanceEvents client: %+v", err)
	}
	configureFunc(maintenanceEventsClient.Client)

	microsoftEntraAdministratorsClient, err := microsoftentraadministrators.NewMicrosoftEntraAdministratorsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MicrosoftEntraAdministrators client: %+v", err)
	}
	configureFunc(microsoftEntraAdministratorsClient.Client)

	microsoftEntraAuthenticationsClient, err := microsoftentraauthentications.NewMicrosoftEntraAuthenticationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MicrosoftEntraAuthentications client: %+v", err)
	}
	configureFunc(microsoftEntraAuthenticationsClient.Client)

	nodesClient, err := nodes.NewNodesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Nodes client: %+v", err)
	}
	configureFunc(nodesClient.Client)

	parameterGroupsClient, err := parametergroups.NewParameterGroupsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ParameterGroups client: %+v", err)
	}
	configureFunc(parameterGroupsClient.Client)

	passwordAuthenticationsClient, err := passwordauthentications.NewPasswordAuthenticationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PasswordAuthentications client: %+v", err)
	}
	configureFunc(passwordAuthenticationsClient.Client)

	poolsClient, err := pools.NewPoolsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Pools client: %+v", err)
	}
	configureFunc(poolsClient.Client)

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

	return &Client{
		Authentications:                 authenticationsClient,
		Capabilities:                    capabilitiesClient,
		ChangeDataCaptureCapabilities:   changeDataCaptureCapabilitiesClient,
		ChangeDataCaptureConfigurations: changeDataCaptureConfigurationsClient,
		ChangeDataCaptureDestinations:   changeDataCaptureDestinationsClient,
		Clusters:                        clustersClient,
		DefaultParameterGroups:          defaultParameterGroupsClient,
		Endpoints:                       endpointsClient,
		FirewallRules:                   firewallRulesClient,
		IdentityAssignments:             identityAssignmentsClient,
		LogCaptureConfigurations:        logCaptureConfigurationsClient,
		LogFiles:                        logFilesClient,
		MaintenanceEvents:               maintenanceEventsClient,
		MicrosoftEntraAdministrators:    microsoftEntraAdministratorsClient,
		MicrosoftEntraAuthentications:   microsoftEntraAuthenticationsClient,
		Nodes:                           nodesClient,
		ParameterGroups:                 parameterGroupsClient,
		PasswordAuthentications:         passwordAuthenticationsClient,
		Pools:                           poolsClient,
		PrivateEndpointConnections:      privateEndpointConnectionsClient,
		PrivateLinkResources:            privateLinkResourcesClient,
	}, nil
}
