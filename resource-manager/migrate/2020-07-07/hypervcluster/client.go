package hypervcluster

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HyperVClusterClient struct {
	Client *resourcemanager.Client
}

func NewHyperVClusterClientWithBaseURI(api environments.Api) (*HyperVClusterClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "hypervcluster", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating HyperVClusterClient: %+v", err)
	}

	return &HyperVClusterClient{
		Client: client,
	}, nil
}
