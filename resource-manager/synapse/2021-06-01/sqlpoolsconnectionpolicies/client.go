package sqlpoolsconnectionpolicies

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolsConnectionPoliciesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewSqlPoolsConnectionPoliciesClientWithBaseURI(endpoint string) SqlPoolsConnectionPoliciesClient {
	return SqlPoolsConnectionPoliciesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
