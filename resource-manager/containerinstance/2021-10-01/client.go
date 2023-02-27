package v2021_10_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerinstance/2021-10-01/containerinstance"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

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
