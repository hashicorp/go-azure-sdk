package gatewayapi

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GatewayApiClient struct {
	Client  autorest.Client
	baseUri string
}

func NewGatewayApiClientWithBaseURI(endpoint string) GatewayApiClient {
	return GatewayApiClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
