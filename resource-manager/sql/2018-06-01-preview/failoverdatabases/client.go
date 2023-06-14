package failoverdatabases

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FailoverDatabasesClient struct {
	Client *resourcemanager.Client
}

func NewFailoverDatabasesClientWithBaseURI(api environments.Api) (*FailoverDatabasesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "failoverdatabases", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating FailoverDatabasesClient: %+v", err)
	}

	return &FailoverDatabasesClient{
		Client: client,
	}, nil
}
