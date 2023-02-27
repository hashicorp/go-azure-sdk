// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2022_06_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/fluidrelay/2022-06-01/fluidrelaycontainers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/fluidrelay/2022-06-01/fluidrelayservers"
)

type Client struct {
	FluidRelayContainers *fluidrelaycontainers.FluidRelayContainersClient
	FluidRelayServers    *fluidrelayservers.FluidRelayServersClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	fluidRelayContainersClient := fluidrelaycontainers.NewFluidRelayContainersClientWithBaseURI(endpoint)
	configureAuthFunc(&fluidRelayContainersClient.Client)

	fluidRelayServersClient := fluidrelayservers.NewFluidRelayServersClientWithBaseURI(endpoint)
	configureAuthFunc(&fluidRelayServersClient.Client)

	return Client{
		FluidRelayContainers: &fluidRelayContainersClient,
		FluidRelayServers:    &fluidRelayServersClient,
	}
}
