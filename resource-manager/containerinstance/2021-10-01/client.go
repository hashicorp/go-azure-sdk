// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2021_10_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerinstance/2021-10-01/containerinstance"
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
