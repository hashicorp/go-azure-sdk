package syncgroups

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SyncGroupsClient struct {
	Client *resourcemanager.Client
}

func NewSyncGroupsClientWithBaseURI(api sdkEnv.Api) (*SyncGroupsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "syncgroups", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SyncGroupsClient: %+v", err)
	}

	return &SyncGroupsClient{
		Client: client,
	}, nil
}
