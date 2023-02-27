// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2022_04_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/applicationinsights/2022-04-01/workbooksapis"
)

type Client struct {
	WorkbooksAPIs *workbooksapis.WorkbooksAPIsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	workbooksAPIsClient := workbooksapis.NewWorkbooksAPIsClientWithBaseURI(endpoint)
	configureAuthFunc(&workbooksAPIsClient.Client)

	return Client{
		WorkbooksAPIs: &workbooksAPIsClient,
	}
}
