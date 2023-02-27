// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2020_11_20

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/applicationinsights/2020-11-20/workbooktemplatesapis"
)

type Client struct {
	WorkbookTemplatesAPIs *workbooktemplatesapis.WorkbookTemplatesAPIsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	workbookTemplatesAPIsClient := workbooktemplatesapis.NewWorkbookTemplatesAPIsClientWithBaseURI(endpoint)
	configureAuthFunc(&workbookTemplatesAPIsClient.Client)

	return Client{
		WorkbookTemplatesAPIs: &workbookTemplatesAPIsClient,
	}
}
