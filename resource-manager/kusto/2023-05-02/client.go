package v2023_05_02

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2023-05-02/attacheddatabaseconfigurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2023-05-02/clusterprincipalassignments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2023-05-02/clusters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2023-05-02/databaseprincipalassignments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2023-05-02/databases"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2023-05-02/dataconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2023-05-02/kusto"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2023-05-02/managedprivateendpoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2023-05-02/outboundnetworkdependenciesendpoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2023-05-02/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2023-05-02/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2023-05-02/scripts"
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
		Scripts:                              scriptsClient,
	}, nil
}
