package migratemysqlazuredbformysqlofflinetasks

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MigrateMySqlAzureDbForMySqlOfflineTasksClient struct {
	Client *resourcemanager.Client
}

func NewMigrateMySqlAzureDbForMySqlOfflineTasksClientWithBaseURI(sdkApi sdkEnv.Api) (*MigrateMySqlAzureDbForMySqlOfflineTasksClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "migratemysqlazuredbformysqlofflinetasks", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MigrateMySqlAzureDbForMySqlOfflineTasksClient: %+v", err)
	}

	return &MigrateMySqlAzureDbForMySqlOfflineTasksClient{
		Client: client,
	}, nil
}
