package v2022_09_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/appplatform/2022-09-01-preview/appplatform"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AppPlatform *appplatform.AppPlatformClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	appPlatformClient, err := appplatform.NewAppPlatformClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AppPlatform client: %+v", err)
	}
	configureFunc(appPlatformClient.Client)

	return &Client{
		AppPlatform: appPlatformClient,
	}, nil
}
