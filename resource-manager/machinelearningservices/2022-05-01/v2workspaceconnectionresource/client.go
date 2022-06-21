package v2workspaceconnectionresource

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type V2WorkspaceConnectionResourceClient struct {
	Client  autorest.Client
	baseUri string
}

func NewV2WorkspaceConnectionResourceClientWithBaseURI(endpoint string) V2WorkspaceConnectionResourceClient {
	return V2WorkspaceConnectionResourceClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
