package mhsmprivateendpointconnections

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MHSMPrivateEndpointConnectionsClient struct {
	Client *resourcemanager.Client
}

func NewMHSMPrivateEndpointConnectionsClientWithBaseURI(sdkApi sdkEnv.Api) (*MHSMPrivateEndpointConnectionsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "mhsmprivateendpointconnections", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MHSMPrivateEndpointConnectionsClient: %+v", err)
	}

	return &MHSMPrivateEndpointConnectionsClient{
		Client: client,
	}, nil
}
