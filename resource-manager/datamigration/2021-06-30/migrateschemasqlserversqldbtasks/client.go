package migrateschemasqlserversqldbtasks

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MigrateSchemaSqlServerSqlDbTasksClient struct {
	Client *resourcemanager.Client
}

func NewMigrateSchemaSqlServerSqlDbTasksClientWithBaseURI(sdkApi sdkEnv.Api) (*MigrateSchemaSqlServerSqlDbTasksClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "migrateschemasqlserversqldbtasks", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MigrateSchemaSqlServerSqlDbTasksClient: %+v", err)
	}

	return &MigrateSchemaSqlServerSqlDbTasksClient{
		Client: client,
	}, nil
}
