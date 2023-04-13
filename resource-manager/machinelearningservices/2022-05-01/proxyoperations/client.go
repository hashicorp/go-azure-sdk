package proxyoperations

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProxyOperationsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewProxyOperationsClientWithBaseURI(endpoint string) ProxyOperationsClient {
	return ProxyOperationsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
