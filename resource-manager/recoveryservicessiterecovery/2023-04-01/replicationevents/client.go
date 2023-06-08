package replicationevents

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReplicationEventsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewReplicationEventsClientWithBaseURI(endpoint string) ReplicationEventsClient {
	return ReplicationEventsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
