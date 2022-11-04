package marketplaceregistrationdefinitions

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MarketplaceRegistrationDefinitionsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewMarketplaceRegistrationDefinitionsClientWithBaseURI(endpoint string) MarketplaceRegistrationDefinitionsClient {
	return MarketplaceRegistrationDefinitionsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
