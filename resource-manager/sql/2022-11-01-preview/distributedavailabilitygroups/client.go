package distributedavailabilitygroups

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DistributedAvailabilityGroupsClient struct {
	Client *resourcemanager.Client
}

func NewDistributedAvailabilityGroupsClientWithBaseURI(api environments.Api) (*DistributedAvailabilityGroupsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "distributedavailabilitygroups", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DistributedAvailabilityGroupsClient: %+v", err)
	}

	return &DistributedAvailabilityGroupsClient{
		Client: client,
	}, nil
}
