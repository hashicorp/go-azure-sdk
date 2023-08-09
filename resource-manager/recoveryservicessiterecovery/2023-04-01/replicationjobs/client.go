package replicationjobs

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReplicationJobsClient struct {
	Client *resourcemanager.Client
}

func NewReplicationJobsClientWithBaseURI(sdkApi sdkEnv.Api) (*ReplicationJobsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "replicationjobs", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ReplicationJobsClient: %+v", err)
	}

	return &ReplicationJobsClient{
		Client: client,
	}, nil
}
