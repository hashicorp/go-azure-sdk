package integrationserviceenvironmentmanagedapi

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntegrationServiceEnvironmentManagedApiClient struct {
	Client *resourcemanager.Client
}

func NewIntegrationServiceEnvironmentManagedApiClientWithBaseURI(api environments.Api) (*IntegrationServiceEnvironmentManagedApiClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "integrationserviceenvironmentmanagedapi", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IntegrationServiceEnvironmentManagedApiClient: %+v", err)
	}

	return &IntegrationServiceEnvironmentManagedApiClient{
		Client: client,
	}, nil
}
