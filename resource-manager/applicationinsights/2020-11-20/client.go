package v2020_11_20

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/applicationinsights/2020-11-20/workbooktemplatesapis"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	WorkbookTemplatesAPIs *workbooktemplatesapis.WorkbookTemplatesAPIsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	workbookTemplatesAPIsClient, err := workbooktemplatesapis.NewWorkbookTemplatesAPIsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkbookTemplatesAPIs client: %+v", err)
	}
	configureFunc(workbookTemplatesAPIsClient.Client)

	return &Client{
		WorkbookTemplatesAPIs: workbookTemplatesAPIsClient,
	}, nil
}
