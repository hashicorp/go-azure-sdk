package migratesqlserversqldbtasks

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MigrateSqlServerSqlDbTasksClient struct {
	Client *resourcemanager.Client
}

func NewMigrateSqlServerSqlDbTasksClientWithBaseURI(sdkApi sdkEnv.Api) (*MigrateSqlServerSqlDbTasksClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "migratesqlserversqldbtasks", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MigrateSqlServerSqlDbTasksClient: %+v", err)
	}

	return &MigrateSqlServerSqlDbTasksClient{
		Client: client,
	}, nil
}
