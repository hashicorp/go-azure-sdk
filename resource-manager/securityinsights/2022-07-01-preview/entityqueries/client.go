package entityqueries

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntityQueriesClient struct {
	Client *resourcemanager.Client
}

func NewEntityQueriesClientWithBaseURI(sdkApi sdkEnv.Api) (*EntityQueriesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "entityqueries", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntityQueriesClient: %+v", err)
	}

	return &EntityQueriesClient{
		Client: client,
	}, nil
}
