package backupslongtermretentionoperations

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackupsLongTermRetentionOperationsClient struct {
	Client *resourcemanager.Client
}

func NewBackupsLongTermRetentionOperationsClientWithBaseURI(sdkApi sdkEnv.Api) (*BackupsLongTermRetentionOperationsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "backupslongtermretentionoperations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating BackupsLongTermRetentionOperationsClient: %+v", err)
	}

	return &BackupsLongTermRetentionOperationsClient{
		Client: client,
	}, nil
}
