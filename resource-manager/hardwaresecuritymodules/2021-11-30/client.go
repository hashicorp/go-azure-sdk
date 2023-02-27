package v2021_11_30

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hardwaresecuritymodules/2021-11-30/dedicatedhsms"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Client struct {
	DedicatedHsms *dedicatedhsms.DedicatedHsmsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	dedicatedHsmsClient := dedicatedhsms.NewDedicatedHsmsClientWithBaseURI(endpoint)
	configureAuthFunc(&dedicatedHsmsClient.Client)

	return Client{
		DedicatedHsms: &dedicatedHsmsClient,
	}
}
