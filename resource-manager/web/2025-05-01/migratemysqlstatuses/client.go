package migratemysqlstatuses

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MigrateMySqlStatusesClient struct {
	Client *resourcemanager.Client
}

func NewMigrateMySqlStatusesClientWithBaseURI(sdkApi sdkEnv.Api) (*MigrateMySqlStatusesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "migratemysqlstatuses", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MigrateMySqlStatusesClient: %+v", err)
	}

	return &MigrateMySqlStatusesClient{
		Client: client,
	}, nil
}
