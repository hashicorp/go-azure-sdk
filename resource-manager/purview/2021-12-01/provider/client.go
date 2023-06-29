package provider

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProviderClient struct {
	Client  autorest.Client
	baseUri string
}

func NewProviderClientWithBaseURI(endpoint string) ProviderClient {
	return ProviderClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
