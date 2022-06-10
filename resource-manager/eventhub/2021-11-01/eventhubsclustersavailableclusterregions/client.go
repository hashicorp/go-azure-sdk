package eventhubsclustersavailableclusterregions

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventHubsClustersAvailableClusterRegionsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewEventHubsClustersAvailableClusterRegionsClientWithBaseURI(endpoint string) EventHubsClustersAvailableClusterRegionsClient {
	return EventHubsClustersAvailableClusterRegionsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
