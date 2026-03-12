package migratemysqlstatusoperationgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MigrateMySqlStatusOperationGroupClient struct {
	Client *resourcemanager.Client
}

func NewMigrateMySqlStatusOperationGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*MigrateMySqlStatusOperationGroupClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "migratemysqlstatusoperationgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MigrateMySqlStatusOperationGroupClient: %+v", err)
	}

	return &MigrateMySqlStatusOperationGroupClient{
		Client: client,
	}, nil
}
