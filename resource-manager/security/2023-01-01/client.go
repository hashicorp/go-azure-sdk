package v2023_01_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2023-01-01/pricings"
)

type Client struct {
	Pricings *pricings.PricingsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	pricingsClient := pricings.NewPricingsClientWithBaseURI(endpoint)
	configureAuthFunc(&pricingsClient.Client)

	return Client{
		Pricings: &pricingsClient,
	}
}
