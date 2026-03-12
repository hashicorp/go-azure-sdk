package vnetconnectiongatewayoperationgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VnetConnectionGatewayOperationGroupClient struct {
	Client *resourcemanager.Client
}

func NewVnetConnectionGatewayOperationGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*VnetConnectionGatewayOperationGroupClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "vnetconnectiongatewayoperationgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VnetConnectionGatewayOperationGroupClient: %+v", err)
	}

	return &VnetConnectionGatewayOperationGroupClient{
		Client: client,
	}, nil
}
