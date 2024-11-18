package v2024_04_13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2024-04-13/attacheddatabaseconfigurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2024-04-13/clusterprincipalassignments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2024-04-13/clusters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2024-04-13/databaseprincipalassignments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2024-04-13/databases"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2024-04-13/dataconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2024-04-13/kusto"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2024-04-13/managedprivateendpoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2024-04-13/outboundnetworkdependenciesendpoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2024-04-13/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2024-04-13/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2024-04-13/sandboxcustomimages"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2024-04-13/scripts"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AttachedDatabaseConfigurations       *attacheddatabaseconfigurations.AttachedDatabaseConfigurationsClient
	ClusterPrincipalAssignments          *clusterprincipalassignments.ClusterPrincipalAssignmentsClient
	Clusters                             *clusters.ClustersClient
	DataConnections                      *dataconnections.DataConnectionsClient
	DatabasePrincipalAssignments         *databaseprincipalassignments.DatabasePrincipalAssignmentsClient
	Databases                            *databases.DatabasesClient
	Kusto                                *kusto.KustoClient
	ManagedPrivateEndpoints              *managedprivateendpoints.ManagedPrivateEndpointsClient
	OutboundNetworkDependenciesEndpoints *outboundnetworkdependenciesendpoints.OutboundNetworkDependenciesEndpointsClient
	PrivateEndpointConnections           *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources                 *privatelinkresources.PrivateLinkResourcesClient
	SandboxCustomImages                  *sandboxcustomimages.SandboxCustomImagesClient
	Scripts                              *scripts.ScriptsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	attachedDatabaseConfigurationsClient, err := attacheddatabaseconfigurations.NewAttachedDatabaseConfigurationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AttachedDatabaseConfigurations client: %+v", err)
	}
	configureFunc(attachedDatabaseConfigurationsClient.Client)

	clusterPrincipalAssignmentsClient, err := clusterprincipalassignments.NewClusterPrincipalAssignmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ClusterPrincipalAssignments client: %+v", err)
	}
	configureFunc(clusterPrincipalAssignmentsClient.Client)

	clustersClient, err := clusters.NewClustersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Clusters client: %+v", err)
	}
	configureFunc(clustersClient.Client)

	dataConnectionsClient, err := dataconnections.NewDataConnectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DataConnections client: %+v", err)
	}
	configureFunc(dataConnectionsClient.Client)

	databasePrincipalAssignmentsClient, err := databaseprincipalassignments.NewDatabasePrincipalAssignmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DatabasePrincipalAssignments client: %+v", err)
	}
	configureFunc(databasePrincipalAssignmentsClient.Client)

	databasesClient, err := databases.NewDatabasesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Databases client: %+v", err)
	}
	configureFunc(databasesClient.Client)

	kustoClient, err := kusto.NewKustoClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Kusto client: %+v", err)
	}
	configureFunc(kustoClient.Client)

	managedPrivateEndpointsClient, err := managedprivateendpoints.NewManagedPrivateEndpointsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedPrivateEndpoints client: %+v", err)
	}
	configureFunc(managedPrivateEndpointsClient.Client)

	outboundNetworkDependenciesEndpointsClient, err := outboundnetworkdependenciesendpoints.NewOutboundNetworkDependenciesEndpointsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OutboundNetworkDependenciesEndpoints client: %+v", err)
	}
	configureFunc(outboundNetworkDependenciesEndpointsClient.Client)

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

	sandboxCustomImagesClient, err := sandboxcustomimages.NewSandboxCustomImagesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SandboxCustomImages client: %+v", err)
	}
	configureFunc(sandboxCustomImagesClient.Client)

	scriptsClient, err := scripts.NewScriptsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Scripts client: %+v", err)
	}
	configureFunc(scriptsClient.Client)

	return &Client{
		AttachedDatabaseConfigurations:       attachedDatabaseConfigurationsClient,
		ClusterPrincipalAssignments:          clusterPrincipalAssignmentsClient,
		Clusters:                             clustersClient,
		DataConnections:                      dataConnectionsClient,
		DatabasePrincipalAssignments:         databasePrincipalAssignmentsClient,
		Databases:                            databasesClient,
		Kusto:                                kustoClient,
		ManagedPrivateEndpoints:              managedPrivateEndpointsClient,
		OutboundNetworkDependenciesEndpoints: outboundNetworkDependenciesEndpointsClient,
		PrivateEndpointConnections:           privateEndpointConnectionsClient,
		PrivateLinkResources:                 privateLinkResourcesClient,
		SandboxCustomImages:                  sandboxCustomImagesClient,
		Scripts:                              scriptsClient,
	}, nil
}
