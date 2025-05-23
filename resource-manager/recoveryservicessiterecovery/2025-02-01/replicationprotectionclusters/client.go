package replicationprotectionclusters

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReplicationProtectionClustersClient struct {
	Client *resourcemanager.Client
}

func NewReplicationProtectionClustersClientWithBaseURI(sdkApi sdkEnv.Api) (*ReplicationProtectionClustersClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "replicationprotectionclusters", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ReplicationProtectionClustersClient: %+v", err)
	}

	return &ReplicationProtectionClustersClient{
		Client: client,
	}, nil
}
