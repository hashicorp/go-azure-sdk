package linkedworkspace

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LinkedWorkspaceClient struct {
	Client  autorest.Client
	baseUri string
}

func NewLinkedWorkspaceClientWithBaseURI(endpoint string) LinkedWorkspaceClient {
	return LinkedWorkspaceClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
