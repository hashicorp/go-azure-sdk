package vnetroutes

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VnetRoutesClient struct {
	Client *resourcemanager.Client
}

func NewVnetRoutesClientWithBaseURI(sdkApi sdkEnv.Api) (*VnetRoutesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "vnetroutes", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VnetRoutesClient: %+v", err)
	}

	return &VnetRoutesClient{
		Client: client,
	}, nil
}
