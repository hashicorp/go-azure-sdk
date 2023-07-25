package staticsites

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StaticSitesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewStaticSitesClientWithBaseURI(endpoint string) StaticSitesClient {
	return StaticSitesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
