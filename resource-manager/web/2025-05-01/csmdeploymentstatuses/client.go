package csmdeploymentstatuses

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CsmDeploymentStatusesClient struct {
	Client *resourcemanager.Client
}

func NewCsmDeploymentStatusesClientWithBaseURI(sdkApi sdkEnv.Api) (*CsmDeploymentStatusesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "csmdeploymentstatuses", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CsmDeploymentStatusesClient: %+v", err)
	}

	return &CsmDeploymentStatusesClient{
		Client: client,
	}, nil
}
