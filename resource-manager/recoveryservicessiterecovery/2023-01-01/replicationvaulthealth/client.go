package replicationvaulthealth

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReplicationVaultHealthClient struct {
	Client *resourcemanager.Client
}

func NewReplicationVaultHealthClientWithBaseURI(sdkApi sdkEnv.Api) (*ReplicationVaultHealthClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "replicationvaulthealth", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ReplicationVaultHealthClient: %+v", err)
	}

	return &ReplicationVaultHealthClient{
		Client: client,
	}, nil
}
