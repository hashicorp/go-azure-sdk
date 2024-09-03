package v2024_08_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2024-08-01/databoundaries"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	DataBoundaries *databoundaries.DataBoundariesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	dataBoundariesClient, err := databoundaries.NewDataBoundariesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DataBoundaries client: %+v", err)
	}
	configureFunc(dataBoundariesClient.Client)

	return &Client{
		DataBoundaries: dataBoundariesClient,
	}, nil
}
