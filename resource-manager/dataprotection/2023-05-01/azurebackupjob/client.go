package azurebackupjob

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureBackupJobClient struct {
	Client *resourcemanager.Client
}

func NewAzureBackupJobClientWithBaseURI(api environments.Api) (*AzureBackupJobClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "azurebackupjob", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AzureBackupJobClient: %+v", err)
	}

	return &AzureBackupJobClient{
		Client: client,
	}, nil
}
