package sqlpoolssecurityalertpolicies

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolsSecurityAlertPoliciesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewSqlPoolsSecurityAlertPoliciesClientWithBaseURI(endpoint string) SqlPoolsSecurityAlertPoliciesClient {
	return SqlPoolsSecurityAlertPoliciesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
