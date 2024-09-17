package sqlpoolstransparentdataencryption

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolsTransparentDataEncryptionClient struct {
	Client *resourcemanager.Client
}

func NewSqlPoolsTransparentDataEncryptionClientWithBaseURI(sdkApi sdkEnv.Api) (*SqlPoolsTransparentDataEncryptionClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "sqlpoolstransparentdataencryption", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SqlPoolsTransparentDataEncryptionClient: %+v", err)
	}

	return &SqlPoolsTransparentDataEncryptionClient{
		Client: client,
	}, nil
}
