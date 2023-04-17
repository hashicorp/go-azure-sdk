package v2023_05_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerinstance/2023-05-01/containerinstance"
)

type Client struct {
	ContainerInstance *containerinstance.ContainerInstanceClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	containerInstanceClient := containerinstance.NewContainerInstanceClientWithBaseURI(endpoint)
	configureAuthFunc(&containerInstanceClient.Client)

	return Client{
		ContainerInstance: &containerInstanceClient,
	}
}
