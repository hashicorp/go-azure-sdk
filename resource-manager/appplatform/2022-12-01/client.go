// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2022_12_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/appplatform/2022-12-01/appplatform"
)

type Client struct {
	AppPlatform *appplatform.AppPlatformClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	appPlatformClient := appplatform.NewAppPlatformClientWithBaseURI(endpoint)
	configureAuthFunc(&appPlatformClient.Client)

	return Client{
		AppPlatform: &appPlatformClient,
	}
}
