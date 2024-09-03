package privateendpointconnection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateEndpointConnectionClient struct {
	Client *resourcemanager.Client
}

func NewPrivateEndpointConnectionClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivateEndpointConnectionClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "privateendpointconnection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivateEndpointConnectionClient: %+v", err)
	}

	return &PrivateEndpointConnectionClient{
		Client: client,
	}, nil
}
