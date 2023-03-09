package replicationmigrationitems

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReplicationMigrationItemsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewReplicationMigrationItemsClientWithBaseURI(endpoint string) ReplicationMigrationItemsClient {
	return ReplicationMigrationItemsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
