package topology

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TopologyClient struct {
	Client  autorest.Client
	baseUri string
}

func NewTopologyClientWithBaseURI(endpoint string) TopologyClient {
	return TopologyClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
