package resourceproviders

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceProvidersClient struct {
	Client  autorest.Client
	baseUri string
}

func NewResourceProvidersClientWithBaseURI(endpoint string) ResourceProvidersClient {
	return ResourceProvidersClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
