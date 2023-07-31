package managedinstanceprivateendpointconnections

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedInstancePrivateEndpointConnectionsClient struct {
	Client *resourcemanager.Client
}

func NewManagedInstancePrivateEndpointConnectionsClientWithBaseURI(api sdkEnv.Api) (*ManagedInstancePrivateEndpointConnectionsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "managedinstanceprivateendpointconnections", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedInstancePrivateEndpointConnectionsClient: %+v", err)
	}

	return &ManagedInstancePrivateEndpointConnectionsClient{
		Client: client,
	}, nil
}
