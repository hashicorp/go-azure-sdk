package replicationprotectionintents

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReplicationProtectionIntentsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewReplicationProtectionIntentsClientWithBaseURI(endpoint string) ReplicationProtectionIntentsClient {
	return ReplicationProtectionIntentsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
