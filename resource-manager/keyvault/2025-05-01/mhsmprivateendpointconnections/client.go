package mhsmprivateendpointconnections

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MhsmPrivateEndpointConnectionsClient struct {
	Client *resourcemanager.Client
}

func NewMhsmPrivateEndpointConnectionsClientWithBaseURI(sdkApi sdkEnv.Api) (*MhsmPrivateEndpointConnectionsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "mhsmprivateendpointconnections", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MhsmPrivateEndpointConnectionsClient: %+v", err)
	}

	return &MhsmPrivateEndpointConnectionsClient{
		Client: client,
	}, nil
}
