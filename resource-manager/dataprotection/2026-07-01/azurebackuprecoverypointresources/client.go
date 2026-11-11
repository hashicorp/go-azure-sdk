package azurebackuprecoverypointresources

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureBackupRecoveryPointResourcesClient struct {
	Client *resourcemanager.Client
}

func NewAzureBackupRecoveryPointResourcesClientWithBaseURI(sdkApi sdkEnv.Api) (*AzureBackupRecoveryPointResourcesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "azurebackuprecoverypointresources", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AzureBackupRecoveryPointResourcesClient: %+v", err)
	}

	return &AzureBackupRecoveryPointResourcesClient{
		Client: client,
	}, nil
}
