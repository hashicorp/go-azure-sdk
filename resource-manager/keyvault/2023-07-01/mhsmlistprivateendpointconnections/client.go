package mhsmlistprivateendpointconnections

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MHSMListPrivateEndpointConnectionsClient struct {
	Client *resourcemanager.Client
}

func NewMHSMListPrivateEndpointConnectionsClientWithBaseURI(sdkApi sdkEnv.Api) (*MHSMListPrivateEndpointConnectionsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "mhsmlistprivateendpointconnections", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MHSMListPrivateEndpointConnectionsClient: %+v", err)
	}

	return &MHSMListPrivateEndpointConnectionsClient{
		Client: client,
	}, nil
}
