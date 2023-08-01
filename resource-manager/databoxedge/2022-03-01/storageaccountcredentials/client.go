package storageaccountcredentials

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StorageAccountCredentialsClient struct {
	Client *resourcemanager.Client
}

func NewStorageAccountCredentialsClientWithBaseURI(sdkApi sdkEnv.Api) (*StorageAccountCredentialsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "storageaccountcredentials", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating StorageAccountCredentialsClient: %+v", err)
	}

	return &StorageAccountCredentialsClient{
		Client: client,
	}, nil
}
