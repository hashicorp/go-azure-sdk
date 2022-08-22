package productsbytag

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProductsByTagClient struct {
	Client  autorest.Client
	baseUri string
}

func NewProductsByTagClientWithBaseURI(endpoint string) ProductsByTagClient {
	return ProductsByTagClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
