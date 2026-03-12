package msdeploystatusslotoperationgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MSDeployStatusSlotOperationGroupClient struct {
	Client *resourcemanager.Client
}

func NewMSDeployStatusSlotOperationGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*MSDeployStatusSlotOperationGroupClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "msdeploystatusslotoperationgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MSDeployStatusSlotOperationGroupClient: %+v", err)
	}

	return &MSDeployStatusSlotOperationGroupClient{
		Client: client,
	}, nil
}
