package v2022_11_08

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresqlhsc/2022-11-08/clusters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresqlhsc/2022-11-08/configurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresqlhsc/2022-11-08/firewallrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresqlhsc/2022-11-08/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresqlhsc/2022-11-08/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresqlhsc/2022-11-08/roles"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresqlhsc/2022-11-08/servers"
)

type Client struct {
	Clusters                   *clusters.ClustersClient
	Configurations             *configurations.ConfigurationsClient
	FirewallRules              *firewallrules.FirewallRulesClient
	PrivateEndpointConnections *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources       *privatelinkresources.PrivateLinkResourcesClient
	Roles                      *roles.RolesClient
	Servers                    *servers.ServersClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	clustersClient := clusters.NewClustersClientWithBaseURI(endpoint)
	configureAuthFunc(&clustersClient.Client)

	configurationsClient := configurations.NewConfigurationsClientWithBaseURI(endpoint)
	configureAuthFunc(&configurationsClient.Client)

	firewallRulesClient := firewallrules.NewFirewallRulesClientWithBaseURI(endpoint)
	configureAuthFunc(&firewallRulesClient.Client)

	privateEndpointConnectionsClient := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&privateLinkResourcesClient.Client)

	rolesClient := roles.NewRolesClientWithBaseURI(endpoint)
	configureAuthFunc(&rolesClient.Client)

	serversClient := servers.NewServersClientWithBaseURI(endpoint)
	configureAuthFunc(&serversClient.Client)

	return Client{
		Clusters:                   &clustersClient,
		Configurations:             &configurationsClient,
		FirewallRules:              &firewallRulesClient,
		PrivateEndpointConnections: &privateEndpointConnectionsClient,
		PrivateLinkResources:       &privateLinkResourcesClient,
		Roles:                      &rolesClient,
		Servers:                    &serversClient,
	}
}
