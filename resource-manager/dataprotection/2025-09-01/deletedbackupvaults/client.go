package deletedbackupvaults

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeletedBackupVaultsClient struct {
	Client *resourcemanager.Client
}

func NewDeletedBackupVaultsClientWithBaseURI(sdkApi sdkEnv.Api) (*DeletedBackupVaultsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "deletedbackupvaults", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeletedBackupVaultsClient: %+v", err)
	}

	return &DeletedBackupVaultsClient{
		Client: client,
	}, nil
}
