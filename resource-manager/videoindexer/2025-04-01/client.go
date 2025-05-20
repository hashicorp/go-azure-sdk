package v2025_04_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/videoindexer/2025-04-01/accesstoken"
	"github.com/hashicorp/go-azure-sdk/resource-manager/videoindexer/2025-04-01/accounts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/videoindexer/2025-04-01/arc"
	"github.com/hashicorp/go-azure-sdk/resource-manager/videoindexer/2025-04-01/nameavailability"
	"github.com/hashicorp/go-azure-sdk/resource-manager/videoindexer/2025-04-01/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/videoindexer/2025-04-01/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AccessToken                *accesstoken.AccessTokenClient
	Accounts                   *accounts.AccountsClient
	Arc                        *arc.ArcClient
	NameAvailability           *nameavailability.NameAvailabilityClient
	PrivateEndpointConnections *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources       *privatelinkresources.PrivateLinkResourcesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	accessTokenClient, err := accesstoken.NewAccessTokenClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AccessToken client: %+v", err)
	}
	configureFunc(accessTokenClient.Client)

	accountsClient, err := accounts.NewAccountsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Accounts client: %+v", err)
	}
	configureFunc(accountsClient.Client)

	arcClient, err := arc.NewArcClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Arc client: %+v", err)
	}
	configureFunc(arcClient.Client)

	nameAvailabilityClient, err := nameavailability.NewNameAvailabilityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NameAvailability client: %+v", err)
	}
	configureFunc(nameAvailabilityClient.Client)

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

	return &Client{
		AccessToken:                accessTokenClient,
		Accounts:                   accountsClient,
		Arc:                        arcClient,
		NameAvailability:           nameAvailabilityClient,
		PrivateEndpointConnections: privateEndpointConnectionsClient,
		PrivateLinkResources:       privateLinkResourcesClient,
	}, nil
}
