package sharedgalleries

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharedGalleriesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewSharedGalleriesClientWithBaseURI(endpoint string) SharedGalleriesClient {
	return SharedGalleriesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
