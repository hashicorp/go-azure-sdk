package v2020_02_02

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/applicationinsights/2020-02-02/componentsapis"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	ComponentsAPIs *componentsapis.ComponentsAPIsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	componentsAPIsClient, err := componentsapis.NewComponentsAPIsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ComponentsAPIs client: %+v", err)
	}
	configureFunc(componentsAPIsClient.Client)

	return &Client{
		ComponentsAPIs: componentsAPIsClient,
	}, nil
}
