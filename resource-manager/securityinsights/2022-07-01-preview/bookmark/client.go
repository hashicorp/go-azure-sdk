package bookmark

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookmarkClient struct {
	Client  autorest.Client
	baseUri string
}

func NewBookmarkClientWithBaseURI(endpoint string) BookmarkClient {
	return BookmarkClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
