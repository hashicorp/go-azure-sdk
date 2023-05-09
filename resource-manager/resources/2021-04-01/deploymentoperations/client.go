package deploymentoperations

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeploymentOperationsClient struct {
	Client *resourcemanager.Client
}

func NewDeploymentOperationsClientWithBaseURI(api environments.Api) (*DeploymentOperationsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "deploymentoperations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeploymentOperationsClient: %+v", err)
	}

	return &DeploymentOperationsClient{
		Client: client,
	}, nil
}
