package backuprestore

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackupRestoreClient struct {
	Client *resourcemanager.Client
}

func NewBackupRestoreClientWithBaseURI(api sdkEnv.Api) (*BackupRestoreClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "backuprestore", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating BackupRestoreClient: %+v", err)
	}

	return &BackupRestoreClient{
		Client: client,
	}, nil
}
