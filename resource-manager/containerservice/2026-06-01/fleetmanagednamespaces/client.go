package fleetmanagednamespaces

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FleetManagedNamespacesClient struct {
	Client *resourcemanager.Client
}

func NewFleetManagedNamespacesClientWithBaseURI(sdkApi sdkEnv.Api) (*FleetManagedNamespacesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "fleetmanagednamespaces", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating FleetManagedNamespacesClient: %+v", err)
	}

	return &FleetManagedNamespacesClient{
		Client: client,
	}, nil
}
