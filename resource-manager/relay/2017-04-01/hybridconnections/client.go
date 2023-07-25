package hybridconnections

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HybridConnectionsClient struct {
	Client *resourcemanager.Client
}

func NewHybridConnectionsClientWithBaseURI(api environments.Api) (*HybridConnectionsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "hybridconnections", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating HybridConnectionsClient: %+v", err)
	}

	return &HybridConnectionsClient{
		Client: client,
	}, nil
}
