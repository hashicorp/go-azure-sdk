package computerps

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComputeRPSClient struct {
	Client *resourcemanager.Client
}

func NewComputeRPSClientWithBaseURI(sdkApi sdkEnv.Api) (*ComputeRPSClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "computerps", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ComputeRPSClient: %+v", err)
	}

	return &ComputeRPSClient{
		Client: client,
	}, nil
}
