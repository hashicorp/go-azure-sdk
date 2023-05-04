package v2020_08_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/search/2020-08-01/adminkeys"
	"github.com/hashicorp/go-azure-sdk/resource-manager/search/2020-08-01/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/search/2020-08-01/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/search/2020-08-01/querykeys"
	"github.com/hashicorp/go-azure-sdk/resource-manager/search/2020-08-01/services"
	"github.com/hashicorp/go-azure-sdk/resource-manager/search/2020-08-01/sharedprivatelinkresources"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AdminKeys                  *adminkeys.AdminKeysClient
	PrivateEndpointConnections *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources       *privatelinkresources.PrivateLinkResourcesClient
	QueryKeys                  *querykeys.QueryKeysClient
	Services                   *services.ServicesClient
	SharedPrivateLinkResources *sharedprivatelinkresources.SharedPrivateLinkResourcesClient
}

func NewClientWithBaseURI(api environments.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	adminKeysClient, err := adminkeys.NewAdminKeysClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building AdminKeys client: %+v", err)
	}
	configureFunc(adminKeysClient.Client)

	privateEndpointConnectionsClient, err := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building PrivateEndpointConnections client: %+v", err)
	}
	configureFunc(privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient, err := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building PrivateLinkResources client: %+v", err)
	}
	configureFunc(privateLinkResourcesClient.Client)

	queryKeysClient, err := querykeys.NewQueryKeysClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building QueryKeys client: %+v", err)
	}
	configureFunc(queryKeysClient.Client)

	servicesClient, err := services.NewServicesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Services client: %+v", err)
	}
	configureFunc(servicesClient.Client)

	sharedPrivateLinkResourcesClient, err := sharedprivatelinkresources.NewSharedPrivateLinkResourcesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building SharedPrivateLinkResources client: %+v", err)
	}
	configureFunc(sharedPrivateLinkResourcesClient.Client)

	return &Client{
		AdminKeys:                  adminKeysClient,
		PrivateEndpointConnections: privateEndpointConnectionsClient,
		PrivateLinkResources:       privateLinkResourcesClient,
		QueryKeys:                  queryKeysClient,
		Services:                   servicesClient,
		SharedPrivateLinkResources: sharedPrivateLinkResourcesClient,
	}, nil
}
