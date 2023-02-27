// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2021_08_15

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/extendedlocation/2021-08-15/customlocations"
)

type Client struct {
	CustomLocations *customlocations.CustomLocationsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	customLocationsClient := customlocations.NewCustomLocationsClientWithBaseURI(endpoint)
	configureAuthFunc(&customLocationsClient.Client)

	return Client{
		CustomLocations: &customLocationsClient,
	}
}
