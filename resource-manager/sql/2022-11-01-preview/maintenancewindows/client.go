package maintenancewindows

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MaintenanceWindowsClient struct {
	Client *resourcemanager.Client
}

func NewMaintenanceWindowsClientWithBaseURI(api environments.Api) (*MaintenanceWindowsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "maintenancewindows", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MaintenanceWindowsClient: %+v", err)
	}

	return &MaintenanceWindowsClient{
		Client: client,
	}, nil
}
