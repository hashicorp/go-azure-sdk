package migratemysqlazuredbformysqlsynctasks

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MigrateMySqlAzureDbForMySqlSyncTasksClient struct {
	Client *resourcemanager.Client
}

func NewMigrateMySqlAzureDbForMySqlSyncTasksClientWithBaseURI(sdkApi sdkEnv.Api) (*MigrateMySqlAzureDbForMySqlSyncTasksClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "migratemysqlazuredbformysqlsynctasks", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MigrateMySqlAzureDbForMySqlSyncTasksClient: %+v", err)
	}

	return &MigrateMySqlAzureDbForMySqlSyncTasksClient{
		Client: client,
	}, nil
}
