package eventhubsclustersconfiguration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventHubsClustersConfigurationClient struct {
	Client *resourcemanager.Client
}

func NewEventHubsClustersConfigurationClientWithBaseURI(sdkApi sdkEnv.Api) (*EventHubsClustersConfigurationClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "eventhubsclustersconfiguration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EventHubsClustersConfigurationClient: %+v", err)
	}

	return &EventHubsClustersConfigurationClient{
		Client: client,
	}, nil
}
