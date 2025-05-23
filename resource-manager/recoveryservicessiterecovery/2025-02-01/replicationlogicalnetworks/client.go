package replicationlogicalnetworks

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReplicationLogicalNetworksClient struct {
	Client *resourcemanager.Client
}

func NewReplicationLogicalNetworksClientWithBaseURI(sdkApi sdkEnv.Api) (*ReplicationLogicalNetworksClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "replicationlogicalnetworks", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ReplicationLogicalNetworksClient: %+v", err)
	}

	return &ReplicationLogicalNetworksClient{
		Client: client,
	}, nil
}
