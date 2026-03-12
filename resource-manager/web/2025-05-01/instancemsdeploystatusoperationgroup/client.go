package instancemsdeploystatusoperationgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InstanceMSDeployStatusOperationGroupClient struct {
	Client *resourcemanager.Client
}

func NewInstanceMSDeployStatusOperationGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*InstanceMSDeployStatusOperationGroupClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "instancemsdeploystatusoperationgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating InstanceMSDeployStatusOperationGroupClient: %+v", err)
	}

	return &InstanceMSDeployStatusOperationGroupClient{
		Client: client,
	}, nil
}
