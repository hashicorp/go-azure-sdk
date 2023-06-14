package manageddatabasecolumns

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDatabaseColumnsClient struct {
	Client *resourcemanager.Client
}

func NewManagedDatabaseColumnsClientWithBaseURI(api environments.Api) (*ManagedDatabaseColumnsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "manageddatabasecolumns", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedDatabaseColumnsClient: %+v", err)
	}

	return &ManagedDatabaseColumnsClient{
		Client: client,
	}, nil
}
