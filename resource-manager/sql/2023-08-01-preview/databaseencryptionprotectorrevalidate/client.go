package databaseencryptionprotectorrevalidate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DatabaseEncryptionProtectorRevalidateClient struct {
	Client *resourcemanager.Client
}

func NewDatabaseEncryptionProtectorRevalidateClientWithBaseURI(sdkApi sdkEnv.Api) (*DatabaseEncryptionProtectorRevalidateClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "databaseencryptionprotectorrevalidate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DatabaseEncryptionProtectorRevalidateClient: %+v", err)
	}

	return &DatabaseEncryptionProtectorRevalidateClient{
		Client: client,
	}, nil
}
