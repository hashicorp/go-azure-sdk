package v2022_03_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2022-03-01/pricings"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Pricings *pricings.PricingsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	pricingsClient, err := pricings.NewPricingsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Pricings client: %+v", err)
	}
	configureFunc(pricingsClient.Client)

	return &Client{
		Pricings: pricingsClient,
	}, nil
}
