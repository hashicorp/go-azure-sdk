package communitygalleryimageversions

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommunityGalleryImageVersionsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewCommunityGalleryImageVersionsClientWithBaseURI(endpoint string) CommunityGalleryImageVersionsClient {
	return CommunityGalleryImageVersionsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
