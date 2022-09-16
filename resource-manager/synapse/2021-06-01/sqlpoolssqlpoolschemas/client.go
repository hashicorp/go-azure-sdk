package sqlpoolssqlpoolschemas

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolsSqlPoolSchemasClient struct {
	Client  autorest.Client
	baseUri string
}

func NewSqlPoolsSqlPoolSchemasClientWithBaseURI(endpoint string) SqlPoolsSqlPoolSchemasClient {
	return SqlPoolsSqlPoolSchemasClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
