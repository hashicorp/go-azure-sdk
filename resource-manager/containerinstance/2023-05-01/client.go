package v2023_05_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/containerinstance/2023-05-01/containerinstance"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	ContainerInstance *containerinstance.ContainerInstanceClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	containerInstanceClient, err := containerinstance.NewContainerInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ContainerInstance client: %+v", err)
	}
	configureFunc(containerInstanceClient.Client)

	return &Client{
		ContainerInstance: containerInstanceClient,
	}, nil
}
