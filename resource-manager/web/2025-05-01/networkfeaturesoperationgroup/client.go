package networkfeaturesoperationgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkFeaturesOperationGroupClient struct {
	Client *resourcemanager.Client
}

func NewNetworkFeaturesOperationGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*NetworkFeaturesOperationGroupClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "networkfeaturesoperationgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating NetworkFeaturesOperationGroupClient: %+v", err)
	}

	return &NetworkFeaturesOperationGroupClient{
		Client: client,
	}, nil
}
