package collection

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CollectionClient struct {
	Client  autorest.Client
	baseUri string
}

func NewCollectionClientWithBaseURI(endpoint string) CollectionClient {
	return CollectionClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
