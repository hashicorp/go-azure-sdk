package restorabledroppedmanageddatabases

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestorableDroppedManagedDatabasesClient struct {
	Client *resourcemanager.Client
}

func NewRestorableDroppedManagedDatabasesClientWithBaseURI(api sdkEnv.Api) (*RestorableDroppedManagedDatabasesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "restorabledroppedmanageddatabases", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RestorableDroppedManagedDatabasesClient: %+v", err)
	}

	return &RestorableDroppedManagedDatabasesClient{
		Client: client,
	}, nil
}
