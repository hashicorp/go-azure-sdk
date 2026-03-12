package backupitemoperationgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackupItemOperationGroupClient struct {
	Client *resourcemanager.Client
}

func NewBackupItemOperationGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*BackupItemOperationGroupClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "backupitemoperationgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating BackupItemOperationGroupClient: %+v", err)
	}

	return &BackupItemOperationGroupClient{
		Client: client,
	}, nil
}
