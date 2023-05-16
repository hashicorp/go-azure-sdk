package datalakestoreaccounts

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataLakeStoreAccountsClient struct {
	Client *resourcemanager.Client
}

func NewDataLakeStoreAccountsClientWithBaseURI(api environments.Api) (*DataLakeStoreAccountsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "datalakestoreaccounts", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DataLakeStoreAccountsClient: %+v", err)
	}

	return &DataLakeStoreAccountsClient{
		Client: client,
	}, nil
}
