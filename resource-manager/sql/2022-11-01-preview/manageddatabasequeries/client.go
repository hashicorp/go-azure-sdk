package manageddatabasequeries

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDatabaseQueriesClient struct {
	Client *resourcemanager.Client
}

func NewManagedDatabaseQueriesClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagedDatabaseQueriesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "manageddatabasequeries", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedDatabaseQueriesClient: %+v", err)
	}

	return &ManagedDatabaseQueriesClient{
		Client: client,
	}, nil
}
