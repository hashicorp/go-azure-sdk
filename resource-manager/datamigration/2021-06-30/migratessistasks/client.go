package migratessistasks

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MigrateSsisTasksClient struct {
	Client *resourcemanager.Client
}

func NewMigrateSsisTasksClientWithBaseURI(sdkApi sdkEnv.Api) (*MigrateSsisTasksClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "migratessistasks", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MigrateSsisTasksClient: %+v", err)
	}

	return &MigrateSsisTasksClient{
		Client: client,
	}, nil
}
