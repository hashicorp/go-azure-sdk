package galleryimageversions

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GalleryImageVersionsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewGalleryImageVersionsClientWithBaseURI(endpoint string) GalleryImageVersionsClient {
	return GalleryImageVersionsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
