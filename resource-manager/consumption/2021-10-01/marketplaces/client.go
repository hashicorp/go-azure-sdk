package marketplaces

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MarketplacesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewMarketplacesClientWithBaseURI(endpoint string) MarketplacesClient {
	return MarketplacesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
