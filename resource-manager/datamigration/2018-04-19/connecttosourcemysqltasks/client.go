package connecttosourcemysqltasks

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectToSourceMySqlTasksClient struct {
	Client *resourcemanager.Client
}

func NewConnectToSourceMySqlTasksClientWithBaseURI(sdkApi sdkEnv.Api) (*ConnectToSourceMySqlTasksClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "connecttosourcemysqltasks", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConnectToSourceMySqlTasksClient: %+v", err)
	}

	return &ConnectToSourceMySqlTasksClient{
		Client: client,
	}, nil
}
