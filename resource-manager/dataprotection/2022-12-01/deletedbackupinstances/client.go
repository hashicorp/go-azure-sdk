package deletedbackupinstances

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeletedBackupInstancesClient struct {
	Client *resourcemanager.Client
}

func NewDeletedBackupInstancesClientWithBaseURI(api environments.Api) (*DeletedBackupInstancesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "deletedbackupinstances", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeletedBackupInstancesClient: %+v", err)
	}

	return &DeletedBackupInstancesClient{
		Client: client,
	}, nil
}
