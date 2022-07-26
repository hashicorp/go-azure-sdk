package namespacesprivatelinkresources

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NamespacesPrivateLinkResourcesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewNamespacesPrivateLinkResourcesClientWithBaseURI(endpoint string) NamespacesPrivateLinkResourcesClient {
	return NamespacesPrivateLinkResourcesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
