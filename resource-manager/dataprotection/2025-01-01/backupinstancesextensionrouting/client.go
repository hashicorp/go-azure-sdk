package backupinstancesextensionrouting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackupInstancesExtensionRoutingClient struct {
	Client *resourcemanager.Client
}

func NewBackupInstancesExtensionRoutingClientWithBaseURI(sdkApi sdkEnv.Api) (*BackupInstancesExtensionRoutingClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "backupinstancesextensionrouting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating BackupInstancesExtensionRoutingClient: %+v", err)
	}

	return &BackupInstancesExtensionRoutingClient{
		Client: client,
	}, nil
}
