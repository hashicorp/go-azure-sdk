package mhsmprivateendpointconnections

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MHSMPrivateEndpointConnectionsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewMHSMPrivateEndpointConnectionsClientWithBaseURI(endpoint string) MHSMPrivateEndpointConnectionsClient {
	return MHSMPrivateEndpointConnectionsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
