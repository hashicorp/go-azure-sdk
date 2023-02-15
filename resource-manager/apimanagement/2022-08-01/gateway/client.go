package gateway

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GatewayClient struct {
	Client  autorest.Client
	baseUri string
}

func NewGatewayClientWithBaseURI(endpoint string) GatewayClient {
	return GatewayClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
