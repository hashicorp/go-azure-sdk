// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2023_01_31

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/managedidentity/2023-01-31/managedidentities"
)

type Client struct {
	ManagedIdentities *managedidentities.ManagedIdentitiesClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	managedIdentitiesClient := managedidentities.NewManagedIdentitiesClientWithBaseURI(endpoint)
	configureAuthFunc(&managedIdentitiesClient.Client)

	return Client{
		ManagedIdentities: &managedIdentitiesClient,
	}
}
