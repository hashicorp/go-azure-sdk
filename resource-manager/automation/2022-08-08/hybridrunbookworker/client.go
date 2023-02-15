package hybridrunbookworker

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HybridRunbookWorkerClient struct {
	Client *resourcemanager.Client
}

func NewHybridRunbookWorkerClientWithBaseURI(api environments.Api) (*HybridRunbookWorkerClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "hybridrunbookworker", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating HybridRunbookWorkerClient: %+v", err)
	}

	return &HybridRunbookWorkerClient{
		Client: client,
	}, nil
}
