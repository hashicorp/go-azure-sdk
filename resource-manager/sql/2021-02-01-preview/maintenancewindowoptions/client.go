package maintenancewindowoptions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MaintenanceWindowOptionsClient struct {
	Client *resourcemanager.Client
}

func NewMaintenanceWindowOptionsClientWithBaseURI(api sdkEnv.Api) (*MaintenanceWindowOptionsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "maintenancewindowoptions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MaintenanceWindowOptionsClient: %+v", err)
	}

	return &MaintenanceWindowOptionsClient{
		Client: client,
	}, nil
}
