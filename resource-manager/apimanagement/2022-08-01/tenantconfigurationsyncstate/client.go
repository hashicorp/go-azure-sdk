package tenantconfigurationsyncstate

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TenantConfigurationSyncStateClient struct {
	Client  autorest.Client
	baseUri string
}

func NewTenantConfigurationSyncStateClientWithBaseURI(endpoint string) TenantConfigurationSyncStateClient {
	return TenantConfigurationSyncStateClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
