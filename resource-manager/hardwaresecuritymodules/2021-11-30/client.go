package v2021_11_30

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/hardwaresecuritymodules/2021-11-30/dedicatedhsms"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	DedicatedHsms *dedicatedhsms.DedicatedHsmsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	dedicatedHsmsClient, err := dedicatedhsms.NewDedicatedHsmsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DedicatedHsms client: %+v", err)
	}
	configureFunc(dedicatedHsmsClient.Client)

	return &Client{
		DedicatedHsms: dedicatedHsmsClient,
	}, nil
}
