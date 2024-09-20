package eventhubsclustersavailableclusterregions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventHubsClustersAvailableClusterRegionsClient struct {
	Client *resourcemanager.Client
}

func NewEventHubsClustersAvailableClusterRegionsClientWithBaseURI(sdkApi sdkEnv.Api) (*EventHubsClustersAvailableClusterRegionsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "eventhubsclustersavailableclusterregions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EventHubsClustersAvailableClusterRegionsClient: %+v", err)
	}

	return &EventHubsClustersAvailableClusterRegionsClient{
		Client: client,
	}, nil
}
