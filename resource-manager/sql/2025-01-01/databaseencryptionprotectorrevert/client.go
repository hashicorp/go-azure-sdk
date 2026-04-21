package databaseencryptionprotectorrevert

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DatabaseEncryptionProtectorRevertClient struct {
	Client *resourcemanager.Client
}

func NewDatabaseEncryptionProtectorRevertClientWithBaseURI(sdkApi sdkEnv.Api) (*DatabaseEncryptionProtectorRevertClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "databaseencryptionprotectorrevert", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DatabaseEncryptionProtectorRevertClient: %+v", err)
	}

	return &DatabaseEncryptionProtectorRevertClient{
		Client: client,
	}, nil
}
