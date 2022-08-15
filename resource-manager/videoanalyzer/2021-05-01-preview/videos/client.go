package videos

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VideosClient struct {
	Client  autorest.Client
	baseUri string
}

func NewVideosClientWithBaseURI(endpoint string) VideosClient {
	return VideosClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
