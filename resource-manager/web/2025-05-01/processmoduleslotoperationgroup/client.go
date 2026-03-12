package processmoduleslotoperationgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProcessModuleSlotOperationGroupClient struct {
	Client *resourcemanager.Client
}

func NewProcessModuleSlotOperationGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*ProcessModuleSlotOperationGroupClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "processmoduleslotoperationgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ProcessModuleSlotOperationGroupClient: %+v", err)
	}

	return &ProcessModuleSlotOperationGroupClient{
		Client: client,
	}, nil
}
