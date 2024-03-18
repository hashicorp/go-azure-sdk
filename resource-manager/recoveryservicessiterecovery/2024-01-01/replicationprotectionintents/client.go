package replicationprotectionintents

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReplicationProtectionIntentsClient struct {
	Client *resourcemanager.Client
}

func NewReplicationProtectionIntentsClientWithBaseURI(sdkApi sdkEnv.Api) (*ReplicationProtectionIntentsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "replicationprotectionintents", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ReplicationProtectionIntentsClient: %+v", err)
	}

	return &ReplicationProtectionIntentsClient{
		Client: client,
	}, nil
}
