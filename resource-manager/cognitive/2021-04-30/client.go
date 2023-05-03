package v2021_04_30

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2021-04-30/cognitiveservicesaccounts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2021-04-30/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2021-04-30/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2021-04-30/skus"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	CognitiveServicesAccounts  *cognitiveservicesaccounts.CognitiveServicesAccountsClient
	PrivateEndpointConnections *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources       *privatelinkresources.PrivateLinkResourcesClient
	Skus                       *skus.SkusClient
}

func NewClientWithBaseURI(api environments.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	cognitiveServicesAccountsClient, err := cognitiveservicesaccounts.NewCognitiveServicesAccountsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building CognitiveServicesAccounts client: %+v", err)
	}
	configureFunc(cognitiveServicesAccountsClient.Client)

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

	skusClient, err := skus.NewSkusClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Skus client: %+v", err)
	}
	configureFunc(skusClient.Client)

	return &Client{
		CognitiveServicesAccounts:  cognitiveServicesAccountsClient,
		PrivateEndpointConnections: privateEndpointConnectionsClient,
		PrivateLinkResources:       privateLinkResourcesClient,
		Skus:                       skusClient,
	}, nil
}
