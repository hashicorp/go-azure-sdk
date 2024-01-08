package contentproductpackages

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentProductPackagesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewContentProductPackagesClientWithBaseURI(endpoint string) ContentProductPackagesClient {
	return ContentProductPackagesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
