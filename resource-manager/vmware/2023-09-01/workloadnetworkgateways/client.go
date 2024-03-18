package workloadnetworkgateways

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkloadNetworkGatewaysClient struct {
	Client *resourcemanager.Client
}

func NewWorkloadNetworkGatewaysClientWithBaseURI(sdkApi sdkEnv.Api) (*WorkloadNetworkGatewaysClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "workloadnetworkgateways", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WorkloadNetworkGatewaysClient: %+v", err)
	}

	return &WorkloadNetworkGatewaysClient{
		Client: client,
	}, nil
}
