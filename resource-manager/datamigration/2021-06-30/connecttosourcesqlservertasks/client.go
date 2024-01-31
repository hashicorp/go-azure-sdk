package connecttosourcesqlservertasks

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectToSourceSqlServerTasksClient struct {
	Client *resourcemanager.Client
}

func NewConnectToSourceSqlServerTasksClientWithBaseURI(sdkApi sdkEnv.Api) (*ConnectToSourceSqlServerTasksClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "connecttosourcesqlservertasks", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConnectToSourceSqlServerTasksClient: %+v", err)
	}

	return &ConnectToSourceSqlServerTasksClient{
		Client: client,
	}, nil
}
