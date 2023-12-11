package v2023_11_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2023-11-01/bicepclient"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	BicepClient *bicepclient.BicepClientClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	bicepClientClient, err := bicepclient.NewBicepClientClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BicepClient client: %+v", err)
	}
	configureFunc(bicepClientClient.Client)

	return &Client{
		BicepClient: bicepClientClient,
	}, nil
}
