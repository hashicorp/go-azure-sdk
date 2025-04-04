package v2025_04_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/redisenterprise/2025-04-01/accesspolicyassignments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/redisenterprise/2025-04-01/databases"
	"github.com/hashicorp/go-azure-sdk/resource-manager/redisenterprise/2025-04-01/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/redisenterprise/2025-04-01/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/redisenterprise/2025-04-01/redisenterprise"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AccessPolicyAssignments    *accesspolicyassignments.AccessPolicyAssignmentsClient
	Databases                  *databases.DatabasesClient
	PrivateEndpointConnections *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources       *privatelinkresources.PrivateLinkResourcesClient
	RedisEnterprise            *redisenterprise.RedisEnterpriseClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	accessPolicyAssignmentsClient, err := accesspolicyassignments.NewAccessPolicyAssignmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AccessPolicyAssignments client: %+v", err)
	}
	configureFunc(accessPolicyAssignmentsClient.Client)

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

	return &Client{
		AccessPolicyAssignments:    accessPolicyAssignmentsClient,
		Databases:                  databasesClient,
		PrivateEndpointConnections: privateEndpointConnectionsClient,
		PrivateLinkResources:       privateLinkResourcesClient,
		RedisEnterprise:            redisEnterpriseClient,
	}, nil
}
