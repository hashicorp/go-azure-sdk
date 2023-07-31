package eventhubsclustersnamespace

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventHubsClustersNamespaceClient struct {
	Client *resourcemanager.Client
}

func NewEventHubsClustersNamespaceClientWithBaseURI(api sdkEnv.Api) (*EventHubsClustersNamespaceClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "eventhubsclustersnamespace", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EventHubsClustersNamespaceClient: %+v", err)
	}

	return &EventHubsClustersNamespaceClient{
		Client: client,
	}, nil
}
