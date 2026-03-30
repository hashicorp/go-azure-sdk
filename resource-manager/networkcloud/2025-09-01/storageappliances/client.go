package storageappliances

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StorageAppliancesClient struct {
	Client *resourcemanager.Client
}

func NewStorageAppliancesClientWithBaseURI(sdkApi sdkEnv.Api) (*StorageAppliancesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "storageappliances", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating StorageAppliancesClient: %+v", err)
	}

	return &StorageAppliancesClient{
		Client: client,
	}, nil
}
