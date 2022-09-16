package privatelinkhubprivatelinkresources

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateLinkHubPrivateLinkResourcesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewPrivateLinkHubPrivateLinkResourcesClientWithBaseURI(endpoint string) PrivateLinkHubPrivateLinkResourcesClient {
	return PrivateLinkHubPrivateLinkResourcesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
