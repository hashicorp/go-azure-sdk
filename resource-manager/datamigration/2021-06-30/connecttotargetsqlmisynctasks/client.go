package connecttotargetsqlmisynctasks

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectToTargetSqlMiSyncTasksClient struct {
	Client *resourcemanager.Client
}

func NewConnectToTargetSqlMiSyncTasksClientWithBaseURI(sdkApi sdkEnv.Api) (*ConnectToTargetSqlMiSyncTasksClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "connecttotargetsqlmisynctasks", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConnectToTargetSqlMiSyncTasksClient: %+v", err)
	}

	return &ConnectToTargetSqlMiSyncTasksClient{
		Client: client,
	}, nil
}
