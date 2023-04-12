package proxy

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProxyClient struct {
	Client  autorest.Client
	baseUri string
}

func NewProxyClientWithBaseURI(endpoint string) ProxyClient {
	return ProxyClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
