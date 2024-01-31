package migratesqlserversqlmitasks

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MigrateSqlServerSqlMITasksClient struct {
	Client *resourcemanager.Client
}

func NewMigrateSqlServerSqlMITasksClientWithBaseURI(sdkApi sdkEnv.Api) (*MigrateSqlServerSqlMITasksClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "migratesqlserversqlmitasks", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MigrateSqlServerSqlMITasksClient: %+v", err)
	}

	return &MigrateSqlServerSqlMITasksClient{
		Client: client,
	}, nil
}
