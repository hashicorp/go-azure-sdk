package sqlpoolsschemas

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolsSchemasClient struct {
	Client  autorest.Client
	baseUri string
}

func NewSqlPoolsSchemasClientWithBaseURI(endpoint string) SqlPoolsSchemasClient {
	return SqlPoolsSchemasClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
