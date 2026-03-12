package vnetinforesourceoperationgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VnetInfoResourceOperationGroupClient struct {
	Client *resourcemanager.Client
}

func NewVnetInfoResourceOperationGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*VnetInfoResourceOperationGroupClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "vnetinforesourceoperationgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VnetInfoResourceOperationGroupClient: %+v", err)
	}

	return &VnetInfoResourceOperationGroupClient{
		Client: client,
	}, nil
}
