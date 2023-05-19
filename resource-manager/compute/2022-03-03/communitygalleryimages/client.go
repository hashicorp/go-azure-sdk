package communitygalleryimages

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommunityGalleryImagesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewCommunityGalleryImagesClientWithBaseURI(endpoint string) CommunityGalleryImagesClient {
	return CommunityGalleryImagesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
