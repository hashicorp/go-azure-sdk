package gatewaylistkeys

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GatewayListKeysClient struct {
	Client  autorest.Client
	baseUri string
}

func NewGatewayListKeysClientWithBaseURI(endpoint string) GatewayListKeysClient {
	return GatewayListKeysClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
