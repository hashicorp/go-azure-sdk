package databaseadvisors

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DatabaseAdvisorsClient struct {
	Client *resourcemanager.Client
}

func NewDatabaseAdvisorsClientWithBaseURI(api environments.Api) (*DatabaseAdvisorsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "databaseadvisors", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DatabaseAdvisorsClient: %+v", err)
	}

	return &DatabaseAdvisorsClient{
		Client: client,
	}, nil
}
