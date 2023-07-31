package v2022_12_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/appplatform/2022-12-01/appplatform"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AppPlatform *appplatform.AppPlatformClient
}

func NewClientWithBaseURI(api sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	appPlatformClient, err := appplatform.NewAppPlatformClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building AppPlatform client: %+v", err)
	}
	configureFunc(appPlatformClient.Client)

	return &Client{
		AppPlatform: appPlatformClient,
	}, nil
}
