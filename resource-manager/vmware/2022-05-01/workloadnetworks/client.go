package workloadnetworks

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkloadNetworksClient struct {
	Client  autorest.Client
	baseUri string
}

func NewWorkloadNetworksClientWithBaseURI(endpoint string) WorkloadNetworksClient {
	return WorkloadNetworksClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
