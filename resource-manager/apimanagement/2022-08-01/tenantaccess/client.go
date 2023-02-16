package tenantaccess

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TenantAccessClient struct {
	Client  autorest.Client
	baseUri string
}

func NewTenantAccessClientWithBaseURI(endpoint string) TenantAccessClient {
	return TenantAccessClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
