package clusterversion

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterVersionClient struct {
	Client  autorest.Client
	baseUri string
}

func NewClusterVersionClientWithBaseURI(endpoint string) ClusterVersionClient {
	return ClusterVersionClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
