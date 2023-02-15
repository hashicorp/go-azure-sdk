package portalconfig

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PortalConfigClient struct {
	Client  autorest.Client
	baseUri string
}

func NewPortalConfigClientWithBaseURI(endpoint string) PortalConfigClient {
	return PortalConfigClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
