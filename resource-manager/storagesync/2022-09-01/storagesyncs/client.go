package storagesyncs

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StoragesyncsClient struct {
	Client *resourcemanager.Client
}

func NewStoragesyncsClientWithBaseURI(sdkApi sdkEnv.Api) (*StoragesyncsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "storagesyncs", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating StoragesyncsClient: %+v", err)
	}

	return &StoragesyncsClient{
		Client: client,
	}, nil
}
