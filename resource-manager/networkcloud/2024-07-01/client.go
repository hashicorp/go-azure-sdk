package v2024_07_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2024-07-01/networkclouds"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Networkclouds *networkclouds.NetworkcloudsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	networkcloudsClient, err := networkclouds.NewNetworkcloudsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Networkclouds client: %+v", err)
	}
	configureFunc(networkcloudsClient.Client)

	return &Client{
		Networkclouds: networkcloudsClient,
	}, nil
}
