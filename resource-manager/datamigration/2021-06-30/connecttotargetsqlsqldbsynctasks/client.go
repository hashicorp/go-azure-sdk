package connecttotargetsqlsqldbsynctasks

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectToTargetSqlSqlDbSyncTasksClient struct {
	Client *resourcemanager.Client
}

func NewConnectToTargetSqlSqlDbSyncTasksClientWithBaseURI(sdkApi sdkEnv.Api) (*ConnectToTargetSqlSqlDbSyncTasksClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "connecttotargetsqlsqldbsynctasks", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConnectToTargetSqlSqlDbSyncTasksClient: %+v", err)
	}

	return &ConnectToTargetSqlSqlDbSyncTasksClient{
		Client: client,
	}, nil
}
