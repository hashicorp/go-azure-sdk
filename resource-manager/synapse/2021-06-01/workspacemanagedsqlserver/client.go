package workspacemanagedsqlserver

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceManagedSqlServerClient struct {
	Client  autorest.Client
	baseUri string
}

func NewWorkspaceManagedSqlServerClientWithBaseURI(endpoint string) WorkspaceManagedSqlServerClient {
	return WorkspaceManagedSqlServerClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
