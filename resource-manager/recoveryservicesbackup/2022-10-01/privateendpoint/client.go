package privateendpoint

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateEndpointClient struct {
	Client  autorest.Client
	baseUri string
}

func NewPrivateEndpointClientWithBaseURI(endpoint string) PrivateEndpointClient {
	return PrivateEndpointClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
