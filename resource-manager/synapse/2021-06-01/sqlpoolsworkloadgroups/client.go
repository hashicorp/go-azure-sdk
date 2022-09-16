package sqlpoolsworkloadgroups

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolsWorkloadGroupsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewSqlPoolsWorkloadGroupsClientWithBaseURI(endpoint string) SqlPoolsWorkloadGroupsClient {
	return SqlPoolsWorkloadGroupsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
