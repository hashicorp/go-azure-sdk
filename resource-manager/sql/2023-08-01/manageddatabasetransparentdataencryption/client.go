package manageddatabasetransparentdataencryption

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDatabaseTransparentDataEncryptionClient struct {
	Client *resourcemanager.Client
}

func NewManagedDatabaseTransparentDataEncryptionClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagedDatabaseTransparentDataEncryptionClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "manageddatabasetransparentdataencryption", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedDatabaseTransparentDataEncryptionClient: %+v", err)
	}

	return &ManagedDatabaseTransparentDataEncryptionClient{
		Client: client,
	}, nil
}
