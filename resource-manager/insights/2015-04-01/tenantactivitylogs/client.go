package tenantactivitylogs

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TenantActivityLogsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewTenantActivityLogsClientWithBaseURI(endpoint string) TenantActivityLogsClient {
	return TenantActivityLogsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
