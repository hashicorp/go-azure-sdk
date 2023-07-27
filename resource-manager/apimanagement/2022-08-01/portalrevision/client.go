package portalrevision

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PortalRevisionClient struct {
	Client  autorest.Client
	baseUri string
}

func NewPortalRevisionClientWithBaseURI(endpoint string) PortalRevisionClient {
	return PortalRevisionClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
