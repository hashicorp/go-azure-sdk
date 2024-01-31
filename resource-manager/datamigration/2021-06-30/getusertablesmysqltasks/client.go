package getusertablesmysqltasks

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetUserTablesMySqlTasksClient struct {
	Client *resourcemanager.Client
}

func NewGetUserTablesMySqlTasksClientWithBaseURI(sdkApi sdkEnv.Api) (*GetUserTablesMySqlTasksClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "getusertablesmysqltasks", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GetUserTablesMySqlTasksClient: %+v", err)
	}

	return &GetUserTablesMySqlTasksClient{
		Client: client,
	}, nil
}
