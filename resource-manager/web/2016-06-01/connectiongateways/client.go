package connectiongateways

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectionGatewaysClient struct {
	Client  autorest.Client
	baseUri string
}

func NewConnectionGatewaysClientWithBaseURI(endpoint string) ConnectionGatewaysClient {
	return ConnectionGatewaysClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
