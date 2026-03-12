package v2023_01_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/storageactions/2023-01-01/storageactions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/storageactions/2023-01-01/storagetasks"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	StorageTasks   *storagetasks.StorageTasksClient
	Storageactions *storageactions.StorageactionsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	storageTasksClient, err := storagetasks.NewStorageTasksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building StorageTasks client: %+v", err)
	}
	configureFunc(storageTasksClient.Client)

	storageactionsClient, err := storageactions.NewStorageactionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Storageactions client: %+v", err)
	}
	configureFunc(storageactionsClient.Client)

	return &Client{
		StorageTasks:   storageTasksClient,
		Storageactions: storageactionsClient,
	}, nil
}
