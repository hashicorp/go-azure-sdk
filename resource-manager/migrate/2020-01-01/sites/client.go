package sites

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SitesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewSitesClientWithBaseURI(endpoint string) SitesClient {
	return SitesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
