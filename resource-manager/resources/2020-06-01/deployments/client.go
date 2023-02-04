package deployments

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeploymentsClient struct {
	Client *resourcemanager.Client
}

func NewDeploymentsClientWithBaseURI(api environments.Api) (*DeploymentsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "deployments", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeploymentsClient: %+v", err)
	}

	return &DeploymentsClient{
		Client: client,
	}, nil
}
