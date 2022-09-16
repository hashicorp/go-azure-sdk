package sqlpoolsworkloadclassifiers

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolsWorkloadClassifiersClient struct {
	Client  autorest.Client
	baseUri string
}

func NewSqlPoolsWorkloadClassifiersClientWithBaseURI(endpoint string) SqlPoolsWorkloadClassifiersClient {
	return SqlPoolsWorkloadClassifiersClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
