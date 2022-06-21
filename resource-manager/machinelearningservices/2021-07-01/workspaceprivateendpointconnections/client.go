package workspaceprivateendpointconnections

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspacePrivateEndpointConnectionsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewWorkspacePrivateEndpointConnectionsClientWithBaseURI(endpoint string) WorkspacePrivateEndpointConnectionsClient {
	return WorkspacePrivateEndpointConnectionsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
