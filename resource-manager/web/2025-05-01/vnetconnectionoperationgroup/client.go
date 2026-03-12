package vnetconnectionoperationgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VnetConnectionOperationGroupClient struct {
	Client *resourcemanager.Client
}

func NewVnetConnectionOperationGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*VnetConnectionOperationGroupClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "vnetconnectionoperationgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VnetConnectionOperationGroupClient: %+v", err)
	}

	return &VnetConnectionOperationGroupClient{
		Client: client,
	}, nil
}
