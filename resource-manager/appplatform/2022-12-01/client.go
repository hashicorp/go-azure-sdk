package v2022_12_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

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
