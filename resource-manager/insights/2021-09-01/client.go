package v2021_09_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2021-09-01/actiongroupsapis"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2021-09-01/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2021-09-01/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2021-09-01/privatelinkscopedresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2021-09-01/privatelinkscopesapis"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	ActionGroupsAPIs           *actiongroupsapis.ActionGroupsAPIsClient
	PrivateEndpointConnections *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources       *privatelinkresources.PrivateLinkResourcesClient
	PrivateLinkScopedResources *privatelinkscopedresources.PrivateLinkScopedResourcesClient
	PrivateLinkScopesAPIs      *privatelinkscopesapis.PrivateLinkScopesAPIsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	actionGroupsAPIsClient, err := actiongroupsapis.NewActionGroupsAPIsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ActionGroupsAPIs client: %+v", err)
	}
	configureFunc(actionGroupsAPIsClient.Client)

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

	privateLinkScopedResourcesClient, err := privatelinkscopedresources.NewPrivateLinkScopedResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateLinkScopedResources client: %+v", err)
	}
	configureFunc(privateLinkScopedResourcesClient.Client)

	privateLinkScopesAPIsClient, err := privatelinkscopesapis.NewPrivateLinkScopesAPIsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateLinkScopesAPIs client: %+v", err)
	}
	configureFunc(privateLinkScopesAPIsClient.Client)

	return &Client{
		ActionGroupsAPIs:           actionGroupsAPIsClient,
		PrivateEndpointConnections: privateEndpointConnectionsClient,
		PrivateLinkResources:       privateLinkResourcesClient,
		PrivateLinkScopedResources: privateLinkScopedResourcesClient,
		PrivateLinkScopesAPIs:      privateLinkScopesAPIsClient,
	}, nil
}
