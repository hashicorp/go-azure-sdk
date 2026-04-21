package manageddatabaseschemas

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDatabaseSchemasClient struct {
	Client *resourcemanager.Client
}

func NewManagedDatabaseSchemasClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagedDatabaseSchemasClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "manageddatabaseschemas", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedDatabaseSchemasClient: %+v", err)
	}

	return &ManagedDatabaseSchemasClient{
		Client: client,
	}, nil
}
