package storageactions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StorageactionsClient struct {
	Client *resourcemanager.Client
}

func NewStorageactionsClientWithBaseURI(sdkApi sdkEnv.Api) (*StorageactionsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "storageactions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating StorageactionsClient: %+v", err)
	}

	return &StorageactionsClient{
		Client: client,
	}, nil
}
