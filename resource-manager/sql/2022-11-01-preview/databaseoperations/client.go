package databaseoperations

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DatabaseOperationsClient struct {
	Client *resourcemanager.Client
}

func NewDatabaseOperationsClientWithBaseURI(api environments.Api) (*DatabaseOperationsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "databaseoperations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DatabaseOperationsClient: %+v", err)
	}

	return &DatabaseOperationsClient{
		Client: client,
	}, nil
}
