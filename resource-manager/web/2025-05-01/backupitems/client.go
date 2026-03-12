package backupitems

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackupItemsClient struct {
	Client *resourcemanager.Client
}

func NewBackupItemsClientWithBaseURI(sdkApi sdkEnv.Api) (*BackupItemsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "backupitems", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating BackupItemsClient: %+v", err)
	}

	return &BackupItemsClient{
		Client: client,
	}, nil
}
