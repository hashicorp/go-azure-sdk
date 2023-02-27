package v2022_04_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/applicationinsights/2022-04-01/workbooksapis"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

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
