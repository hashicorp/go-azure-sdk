package resourcegroups

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceGroupsClient struct {
	Client *resourcemanager.Client
}

func NewResourceGroupsClientWithBaseURI(api environments.Api) (*ResourceGroupsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "resourcegroups", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ResourceGroupsClient: %+v", err)
	}

	return &ResourceGroupsClient{
		Client: client,
	}, nil
}
