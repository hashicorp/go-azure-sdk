package workspaceskus

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceSkusClient struct {
	Client  autorest.Client
	baseUri string
}

func NewWorkspaceSkusClientWithBaseURI(endpoint string) WorkspaceSkusClient {
	return WorkspaceSkusClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
