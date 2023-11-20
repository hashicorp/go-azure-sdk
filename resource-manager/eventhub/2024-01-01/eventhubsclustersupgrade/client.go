package eventhubsclustersupgrade

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventHubsClustersUpgradeClient struct {
	Client *resourcemanager.Client
}

func NewEventHubsClustersUpgradeClientWithBaseURI(sdkApi sdkEnv.Api) (*EventHubsClustersUpgradeClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "eventhubsclustersupgrade", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EventHubsClustersUpgradeClient: %+v", err)
	}

	return &EventHubsClustersUpgradeClient{
		Client: client,
	}, nil
}
