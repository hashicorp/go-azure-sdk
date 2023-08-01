package storagesyncservice

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StorageSyncServiceClient struct {
	Client *resourcemanager.Client
}

func NewStorageSyncServiceClientWithBaseURI(sdkApi sdkEnv.Api) (*StorageSyncServiceClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "storagesyncservice", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating StorageSyncServiceClient: %+v", err)
	}

	return &StorageSyncServiceClient{
		Client: client,
	}, nil
}
