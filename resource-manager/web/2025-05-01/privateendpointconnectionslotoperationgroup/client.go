package privateendpointconnectionslotoperationgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateEndpointConnectionSlotOperationGroupClient struct {
	Client *resourcemanager.Client
}

func NewPrivateEndpointConnectionSlotOperationGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivateEndpointConnectionSlotOperationGroupClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "privateendpointconnectionslotoperationgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivateEndpointConnectionSlotOperationGroupClient: %+v", err)
	}

	return &PrivateEndpointConnectionSlotOperationGroupClient{
		Client: client,
	}, nil
}
