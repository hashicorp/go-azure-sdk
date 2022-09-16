package sqlpoolsdatamaskingpolicies

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolsDataMaskingPoliciesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewSqlPoolsDataMaskingPoliciesClientWithBaseURI(endpoint string) SqlPoolsDataMaskingPoliciesClient {
	return SqlPoolsDataMaskingPoliciesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
