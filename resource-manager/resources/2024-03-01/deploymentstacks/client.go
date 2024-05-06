package deploymentstacks

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeploymentStacksClient struct {
	Client *resourcemanager.Client
}

func NewDeploymentStacksClientWithBaseURI(sdkApi sdkEnv.Api) (*DeploymentStacksClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "deploymentstacks", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeploymentStacksClient: %+v", err)
	}

	return &DeploymentStacksClient{
		Client: client,
	}, nil
}
