package v2022_04_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/applicationinsights/2022-04-01/workbooksapis"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	WorkbooksAPIs *workbooksapis.WorkbooksAPIsClient
}

func NewClientWithBaseURI(api environments.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	workbooksAPIsClient, err := workbooksapis.NewWorkbooksAPIsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building WorkbooksAPIs client: %+v", err)
	}
	configureFunc(workbooksAPIsClient.Client)

	return &Client{
		WorkbooksAPIs: workbooksAPIsClient,
	}, nil
}
