package workspaceprivatelinkresources

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspacePrivateLinkResourcesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewWorkspacePrivateLinkResourcesClientWithBaseURI(endpoint string) WorkspacePrivateLinkResourcesClient {
	return WorkspacePrivateLinkResourcesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
