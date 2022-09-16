package sqlpoolsblobauditing

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolsBlobAuditingClient struct {
	Client  autorest.Client
	baseUri string
}

func NewSqlPoolsBlobAuditingClientWithBaseURI(endpoint string) SqlPoolsBlobAuditingClient {
	return SqlPoolsBlobAuditingClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
