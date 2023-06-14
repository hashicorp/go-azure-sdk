package syncmembers

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SyncMembersClient struct {
	Client *resourcemanager.Client
}

func NewSyncMembersClientWithBaseURI(api environments.Api) (*SyncMembersClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "syncmembers", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SyncMembersClient: %+v", err)
	}

	return &SyncMembersClient{
		Client: client,
	}, nil
}
