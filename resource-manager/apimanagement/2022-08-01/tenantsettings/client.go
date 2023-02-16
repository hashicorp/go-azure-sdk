package tenantsettings

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TenantSettingsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewTenantSettingsClientWithBaseURI(endpoint string) TenantSettingsClient {
	return TenantSettingsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
