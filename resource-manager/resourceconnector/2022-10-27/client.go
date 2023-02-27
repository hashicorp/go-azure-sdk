// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2022_10_27

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resourceconnector/2022-10-27/appliances"
)

type Client struct {
	Appliances *appliances.AppliancesClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	appliancesClient := appliances.NewAppliancesClientWithBaseURI(endpoint)
	configureAuthFunc(&appliancesClient.Client)

	return Client{
		Appliances: &appliancesClient,
	}
}
