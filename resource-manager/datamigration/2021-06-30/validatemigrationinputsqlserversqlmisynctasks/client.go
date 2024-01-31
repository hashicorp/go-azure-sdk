package validatemigrationinputsqlserversqlmisynctasks

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ValidateMigrationInputSqlServerSqlMiSyncTasksClient struct {
	Client *resourcemanager.Client
}

func NewValidateMigrationInputSqlServerSqlMiSyncTasksClientWithBaseURI(sdkApi sdkEnv.Api) (*ValidateMigrationInputSqlServerSqlMiSyncTasksClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "validatemigrationinputsqlserversqlmisynctasks", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ValidateMigrationInputSqlServerSqlMiSyncTasksClient: %+v", err)
	}

	return &ValidateMigrationInputSqlServerSqlMiSyncTasksClient{
		Client: client,
	}, nil
}
