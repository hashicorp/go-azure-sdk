package listprivatelinkresources

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListPrivateLinkResourcesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewListPrivateLinkResourcesClientWithBaseURI(endpoint string) ListPrivateLinkResourcesClient {
	return ListPrivateLinkResourcesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
