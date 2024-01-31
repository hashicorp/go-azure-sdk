package getusertablessqlsynctasks

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetUserTablesSqlSyncTasksClient struct {
	Client *resourcemanager.Client
}

func NewGetUserTablesSqlSyncTasksClientWithBaseURI(sdkApi sdkEnv.Api) (*GetUserTablesSqlSyncTasksClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "getusertablessqlsynctasks", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GetUserTablesSqlSyncTasksClient: %+v", err)
	}

	return &GetUserTablesSqlSyncTasksClient{
		Client: client,
	}, nil
}
