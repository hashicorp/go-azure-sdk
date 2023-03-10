package v2022_10_27

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

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
