package workspacemanagedsqlserversqlusages

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceManagedSqlServerSqlUsagesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewWorkspaceManagedSqlServerSqlUsagesClientWithBaseURI(endpoint string) WorkspaceManagedSqlServerSqlUsagesClient {
	return WorkspaceManagedSqlServerSqlUsagesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
