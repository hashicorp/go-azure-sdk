// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2022_09_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerinstance/2022-09-01/containerinstance"
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
