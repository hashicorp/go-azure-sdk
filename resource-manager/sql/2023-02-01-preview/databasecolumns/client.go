package databasecolumns

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DatabaseColumnsClient struct {
	Client *resourcemanager.Client
}

func NewDatabaseColumnsClientWithBaseURI(sdkApi sdkEnv.Api) (*DatabaseColumnsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "databasecolumns", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DatabaseColumnsClient: %+v", err)
	}

	return &DatabaseColumnsClient{
		Client: client,
	}, nil
}
