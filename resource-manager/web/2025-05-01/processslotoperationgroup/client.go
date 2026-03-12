package processslotoperationgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProcessSlotOperationGroupClient struct {
	Client *resourcemanager.Client
}

func NewProcessSlotOperationGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*ProcessSlotOperationGroupClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "processslotoperationgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ProcessSlotOperationGroupClient: %+v", err)
	}

	return &ProcessSlotOperationGroupClient{
		Client: client,
	}, nil
}
