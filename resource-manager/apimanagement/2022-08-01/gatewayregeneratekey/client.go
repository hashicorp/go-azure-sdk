package gatewayregeneratekey

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GatewayRegenerateKeyClient struct {
	Client  autorest.Client
	baseUri string
}

func NewGatewayRegenerateKeyClientWithBaseURI(endpoint string) GatewayRegenerateKeyClient {
	return GatewayRegenerateKeyClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
