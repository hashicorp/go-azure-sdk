package sqlpoolsreplicationlinks

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolsReplicationLinksClient struct {
	Client  autorest.Client
	baseUri string
}

func NewSqlPoolsReplicationLinksClientWithBaseURI(endpoint string) SqlPoolsReplicationLinksClient {
	return SqlPoolsReplicationLinksClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
