package workloadnetworks

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkloadNetworksClient struct {
	Client *resourcemanager.Client
}

func NewWorkloadNetworksClientWithBaseURI(sdkApi sdkEnv.Api) (*WorkloadNetworksClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "workloadnetworks", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WorkloadNetworksClient: %+v", err)
	}

	return &WorkloadNetworksClient{
		Client: client,
	}, nil
}
