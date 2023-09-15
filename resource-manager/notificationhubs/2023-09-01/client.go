package v2023_09_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/notificationhubs/2023-09-01/hubs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/notificationhubs/2023-09-01/namespaces"
	"github.com/hashicorp/go-azure-sdk/resource-manager/notificationhubs/2023-09-01/privatelink"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Hubs        *hubs.HubsClient
	Namespaces  *namespaces.NamespacesClient
	PrivateLink *privatelink.PrivateLinkClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	hubsClient, err := hubs.NewHubsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Hubs client: %+v", err)
	}
	configureFunc(hubsClient.Client)

	namespacesClient, err := namespaces.NewNamespacesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Namespaces client: %+v", err)
	}
	configureFunc(namespacesClient.Client)

	privateLinkClient, err := privatelink.NewPrivateLinkClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateLink client: %+v", err)
	}
	configureFunc(privateLinkClient.Client)

	return &Client{
		Hubs:        hubsClient,
		Namespaces:  namespacesClient,
		PrivateLink: privateLinkClient,
	}, nil
}
