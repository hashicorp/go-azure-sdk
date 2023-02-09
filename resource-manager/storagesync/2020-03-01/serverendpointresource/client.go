package serverendpointresource

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerEndpointResourceClient struct {
	Client  autorest.Client
	baseUri string
}

func NewServerEndpointResourceClientWithBaseURI(endpoint string) ServerEndpointResourceClient {
	return ServerEndpointResourceClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
