package bookmarks

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookmarksClient struct {
	Client  autorest.Client
	baseUri string
}

func NewBookmarksClientWithBaseURI(endpoint string) BookmarksClient {
	return BookmarksClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
