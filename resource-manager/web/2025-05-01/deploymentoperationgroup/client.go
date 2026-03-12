package deploymentoperationgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeploymentOperationGroupClient struct {
	Client *resourcemanager.Client
}

func NewDeploymentOperationGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*DeploymentOperationGroupClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "deploymentoperationgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeploymentOperationGroupClient: %+v", err)
	}

	return &DeploymentOperationGroupClient{
		Client: client,
	}, nil
}
