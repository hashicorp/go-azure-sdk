package replicationmigrationitems

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReplicationMigrationItemsClient struct {
	Client *resourcemanager.Client
}

func NewReplicationMigrationItemsClientWithBaseURI(sdkApi sdkEnv.Api) (*ReplicationMigrationItemsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "replicationmigrationitems", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ReplicationMigrationItemsClient: %+v", err)
	}

	return &ReplicationMigrationItemsClient{
		Client: client,
	}, nil
}
