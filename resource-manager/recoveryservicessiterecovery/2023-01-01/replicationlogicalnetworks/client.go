package replicationlogicalnetworks

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReplicationLogicalNetworksClient struct {
	Client  autorest.Client
	baseUri string
}

func NewReplicationLogicalNetworksClientWithBaseURI(endpoint string) ReplicationLogicalNetworksClient {
	return ReplicationLogicalNetworksClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
