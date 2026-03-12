package csmdeploymentstatusoperationgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CsmDeploymentStatusOperationGroupClient struct {
	Client *resourcemanager.Client
}

func NewCsmDeploymentStatusOperationGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*CsmDeploymentStatusOperationGroupClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "csmdeploymentstatusoperationgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CsmDeploymentStatusOperationGroupClient: %+v", err)
	}

	return &CsmDeploymentStatusOperationGroupClient{
		Client: client,
	}, nil
}
