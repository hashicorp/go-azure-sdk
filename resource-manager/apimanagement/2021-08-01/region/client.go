package region

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegionClient struct {
	Client *resourcemanager.Client
}

func NewRegionClientWithBaseURI(api environments.Api) (*RegionClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "region", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RegionClient: %+v", err)
	}

	return &RegionClient{
		Client: client,
	}, nil
}
