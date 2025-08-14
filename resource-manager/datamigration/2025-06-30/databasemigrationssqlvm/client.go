package databasemigrationssqlvm

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DatabaseMigrationsSqlVMClient struct {
	Client *resourcemanager.Client
}

func NewDatabaseMigrationsSqlVMClientWithBaseURI(sdkApi sdkEnv.Api) (*DatabaseMigrationsSqlVMClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "databasemigrationssqlvm", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DatabaseMigrationsSqlVMClient: %+v", err)
	}

	return &DatabaseMigrationsSqlVMClient{
		Client: client,
	}, nil
}
