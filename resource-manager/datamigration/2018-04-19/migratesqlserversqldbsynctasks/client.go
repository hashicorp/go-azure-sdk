package migratesqlserversqldbsynctasks

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MigrateSqlServerSqlDbSyncTasksClient struct {
	Client *resourcemanager.Client
}

func NewMigrateSqlServerSqlDbSyncTasksClientWithBaseURI(sdkApi sdkEnv.Api) (*MigrateSqlServerSqlDbSyncTasksClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "migratesqlserversqldbsynctasks", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MigrateSqlServerSqlDbSyncTasksClient: %+v", err)
	}

	return &MigrateSqlServerSqlDbSyncTasksClient{
		Client: client,
	}, nil
}
