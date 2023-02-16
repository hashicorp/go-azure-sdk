package productwiki

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProductWikiClient struct {
	Client  autorest.Client
	baseUri string
}

func NewProductWikiClientWithBaseURI(endpoint string) ProductWikiClient {
	return ProductWikiClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
