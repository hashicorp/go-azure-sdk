package replicationvaulthealth

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReplicationVaultHealthClient struct {
	Client  autorest.Client
	baseUri string
}

func NewReplicationVaultHealthClientWithBaseURI(endpoint string) ReplicationVaultHealthClient {
	return ReplicationVaultHealthClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
