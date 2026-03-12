package networkfeaturesslotoperationgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkFeaturesSlotOperationGroupClient struct {
	Client *resourcemanager.Client
}

func NewNetworkFeaturesSlotOperationGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*NetworkFeaturesSlotOperationGroupClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "networkfeaturesslotoperationgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating NetworkFeaturesSlotOperationGroupClient: %+v", err)
	}

	return &NetworkFeaturesSlotOperationGroupClient{
		Client: client,
	}, nil
}
