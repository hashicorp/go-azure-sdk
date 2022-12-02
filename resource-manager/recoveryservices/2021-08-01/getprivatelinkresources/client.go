package getprivatelinkresources

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetPrivateLinkResourcesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewGetPrivateLinkResourcesClientWithBaseURI(endpoint string) GetPrivateLinkResourcesClient {
	return GetPrivateLinkResourcesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
