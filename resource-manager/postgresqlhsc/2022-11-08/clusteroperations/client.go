package clusteroperations

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterOperationsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewClusterOperationsClientWithBaseURI(endpoint string) ClusterOperationsClient {
	return ClusterOperationsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
