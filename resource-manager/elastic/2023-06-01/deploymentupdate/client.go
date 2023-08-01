package deploymentupdate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeploymentUpdateClient struct {
	Client *resourcemanager.Client
}

func NewDeploymentUpdateClientWithBaseURI(sdkApi sdkEnv.Api) (*DeploymentUpdateClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "deploymentupdate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeploymentUpdateClient: %+v", err)
	}

	return &DeploymentUpdateClient{
		Client: client,
	}, nil
}
