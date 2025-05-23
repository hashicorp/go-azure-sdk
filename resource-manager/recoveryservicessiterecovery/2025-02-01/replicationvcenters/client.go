package replicationvcenters

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReplicationvCentersClient struct {
	Client *resourcemanager.Client
}

func NewReplicationvCentersClientWithBaseURI(sdkApi sdkEnv.Api) (*ReplicationvCentersClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "replicationvcenters", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ReplicationvCentersClient: %+v", err)
	}

	return &ReplicationvCentersClient{
		Client: client,
	}, nil
}
