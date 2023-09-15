package d4apicollectionlist

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type D4APICollectionListClient struct {
	Client  autorest.Client
	baseUri string
}

func NewD4APICollectionListClientWithBaseURI(endpoint string) D4APICollectionListClient {
	return D4APICollectionListClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
