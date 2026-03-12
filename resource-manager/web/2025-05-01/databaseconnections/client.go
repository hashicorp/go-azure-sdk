package databaseconnections

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DatabaseConnectionsClient struct {
	Client *resourcemanager.Client
}

func NewDatabaseConnectionsClientWithBaseURI(sdkApi sdkEnv.Api) (*DatabaseConnectionsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "databaseconnections", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DatabaseConnectionsClient: %+v", err)
	}

	return &DatabaseConnectionsClient{
		Client: client,
	}, nil
}
