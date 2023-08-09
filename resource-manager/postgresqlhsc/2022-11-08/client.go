package v2022_11_08

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresqlhsc/2022-11-08/clusteroperations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresqlhsc/2022-11-08/clusters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresqlhsc/2022-11-08/configurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresqlhsc/2022-11-08/firewallrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresqlhsc/2022-11-08/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresqlhsc/2022-11-08/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresqlhsc/2022-11-08/roles"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresqlhsc/2022-11-08/servers"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	ClusterOperations          *clusteroperations.ClusterOperationsClient
	Clusters                   *clusters.ClustersClient
	Configurations             *configurations.ConfigurationsClient
	FirewallRules              *firewallrules.FirewallRulesClient
	PrivateEndpointConnections *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources       *privatelinkresources.PrivateLinkResourcesClient
	Roles                      *roles.RolesClient
	Servers                    *servers.ServersClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	clusterOperationsClient, err := clusteroperations.NewClusterOperationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ClusterOperations client: %+v", err)
	}
	configureFunc(clusterOperationsClient.Client)

	clustersClient, err := clusters.NewClustersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Clusters client: %+v", err)
	}
	configureFunc(clustersClient.Client)

	configurationsClient, err := configurations.NewConfigurationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Configurations client: %+v", err)
	}
	configureFunc(configurationsClient.Client)

	firewallRulesClient, err := firewallrules.NewFirewallRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FirewallRules client: %+v", err)
	}
	configureFunc(firewallRulesClient.Client)

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

	rolesClient, err := roles.NewRolesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Roles client: %+v", err)
	}
	configureFunc(rolesClient.Client)

	serversClient, err := servers.NewServersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Servers client: %+v", err)
	}
	configureFunc(serversClient.Client)

	return &Client{
		ClusterOperations:          clusterOperationsClient,
		Clusters:                   clustersClient,
		Configurations:             configurationsClient,
		FirewallRules:              firewallRulesClient,
		PrivateEndpointConnections: privateEndpointConnectionsClient,
		PrivateLinkResources:       privateLinkResourcesClient,
		Roles:                      rolesClient,
		Servers:                    serversClient,
	}, nil
}
