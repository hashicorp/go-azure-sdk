package v2020_03_13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/search/2020-03-13/adminkeys"
	"github.com/hashicorp/go-azure-sdk/resource-manager/search/2020-03-13/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/search/2020-03-13/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/search/2020-03-13/querykeys"
	"github.com/hashicorp/go-azure-sdk/resource-manager/search/2020-03-13/services"
)

type Client struct {
	AdminKeys                  *adminkeys.AdminKeysClient
	PrivateEndpointConnections *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources       *privatelinkresources.PrivateLinkResourcesClient
	QueryKeys                  *querykeys.QueryKeysClient
	Services                   *services.ServicesClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	adminKeysClient := adminkeys.NewAdminKeysClientWithBaseURI(endpoint)
	configureAuthFunc(&adminKeysClient.Client)

	privateEndpointConnectionsClient := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&privateLinkResourcesClient.Client)

	queryKeysClient := querykeys.NewQueryKeysClientWithBaseURI(endpoint)
	configureAuthFunc(&queryKeysClient.Client)

	servicesClient := services.NewServicesClientWithBaseURI(endpoint)
	configureAuthFunc(&servicesClient.Client)

	return Client{
		AdminKeys:                  &adminKeysClient,
		PrivateEndpointConnections: &privateEndpointConnectionsClient,
		PrivateLinkResources:       &privateLinkResourcesClient,
		QueryKeys:                  &queryKeysClient,
		Services:                   &servicesClient,
	}
}
