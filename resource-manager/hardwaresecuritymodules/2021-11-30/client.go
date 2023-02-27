// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2021_11_30

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hardwaresecuritymodules/2021-11-30/dedicatedhsms"
)

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
