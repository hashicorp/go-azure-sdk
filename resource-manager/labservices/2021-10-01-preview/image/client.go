package image

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImageClient struct {
	Client  autorest.Client
	baseUri string
}

func NewImageClientWithBaseURI(endpoint string) ImageClient {
	return ImageClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
