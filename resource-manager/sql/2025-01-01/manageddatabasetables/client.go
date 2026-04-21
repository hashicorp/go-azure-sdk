package manageddatabasetables

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDatabaseTablesClient struct {
	Client *resourcemanager.Client
}

func NewManagedDatabaseTablesClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagedDatabaseTablesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "manageddatabasetables", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedDatabaseTablesClient: %+v", err)
	}

	return &ManagedDatabaseTablesClient{
		Client: client,
	}, nil
}
