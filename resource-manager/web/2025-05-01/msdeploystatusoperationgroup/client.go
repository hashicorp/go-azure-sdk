package msdeploystatusoperationgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MSDeployStatusOperationGroupClient struct {
	Client *resourcemanager.Client
}

func NewMSDeployStatusOperationGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*MSDeployStatusOperationGroupClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "msdeploystatusoperationgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MSDeployStatusOperationGroupClient: %+v", err)
	}

	return &MSDeployStatusOperationGroupClient{
		Client: client,
	}, nil
}
