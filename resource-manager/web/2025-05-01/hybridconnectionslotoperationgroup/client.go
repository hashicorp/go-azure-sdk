package hybridconnectionslotoperationgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HybridConnectionSlotOperationGroupClient struct {
	Client *resourcemanager.Client
}

func NewHybridConnectionSlotOperationGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*HybridConnectionSlotOperationGroupClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "hybridconnectionslotoperationgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating HybridConnectionSlotOperationGroupClient: %+v", err)
	}

	return &HybridConnectionSlotOperationGroupClient{
		Client: client,
	}, nil
}
