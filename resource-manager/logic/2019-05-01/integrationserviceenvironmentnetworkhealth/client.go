package integrationserviceenvironmentnetworkhealth

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntegrationServiceEnvironmentNetworkHealthClient struct {
	Client *resourcemanager.Client
}

func NewIntegrationServiceEnvironmentNetworkHealthClientWithBaseURI(api sdkEnv.Api) (*IntegrationServiceEnvironmentNetworkHealthClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "integrationserviceenvironmentnetworkhealth", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IntegrationServiceEnvironmentNetworkHealthClient: %+v", err)
	}

	return &IntegrationServiceEnvironmentNetworkHealthClient{
		Client: client,
	}, nil
}
