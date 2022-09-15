package sharedgalleryimageversions

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharedGalleryImageVersionsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewSharedGalleryImageVersionsClientWithBaseURI(endpoint string) SharedGalleryImageVersionsClient {
	return SharedGalleryImageVersionsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
