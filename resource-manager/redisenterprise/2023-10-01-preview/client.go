package v2023_10_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/redisenterprise/2023-10-01-preview/databases"
	"github.com/hashicorp/go-azure-sdk/resource-manager/redisenterprise/2023-10-01-preview/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/redisenterprise/2023-10-01-preview/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/redisenterprise/2023-10-01-preview/redisenterprise"
	"github.com/hashicorp/go-azure-sdk/resource-manager/redisenterprise/2023-10-01-preview/sku"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Databases                  *databases.DatabasesClient
	PrivateEndpointConnections *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources       *privatelinkresources.PrivateLinkResourcesClient
	RedisEnterprise            *redisenterprise.RedisEnterpriseClient
	SKU                        *sku.SKUClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	databasesClient, err := databases.NewDatabasesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Databases client: %+v", err)
	}
	configureFunc(databasesClient.Client)

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

	redisEnterpriseClient, err := redisenterprise.NewRedisEnterpriseClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RedisEnterprise client: %+v", err)
	}
	configureFunc(redisEnterpriseClient.Client)

	sKUClient, err := sku.NewSKUClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SKU client: %+v", err)
	}
	configureFunc(sKUClient.Client)

	return &Client{
		Databases:                  databasesClient,
		PrivateEndpointConnections: privateEndpointConnectionsClient,
		PrivateLinkResources:       privateLinkResourcesClient,
		RedisEnterprise:            redisEnterpriseClient,
		SKU:                        sKUClient,
	}, nil
}
