package namespacesprivateendpointconnections

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NamespacesPrivateEndpointConnectionsClient struct {
	Client *resourcemanager.Client
}

func NewNamespacesPrivateEndpointConnectionsClientWithBaseURI(api sdkEnv.Api) (*NamespacesPrivateEndpointConnectionsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "namespacesprivateendpointconnections", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating NamespacesPrivateEndpointConnectionsClient: %+v", err)
	}

	return &NamespacesPrivateEndpointConnectionsClient{
		Client: client,
	}, nil
}
