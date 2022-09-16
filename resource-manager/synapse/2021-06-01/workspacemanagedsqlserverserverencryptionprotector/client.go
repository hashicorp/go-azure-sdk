package workspacemanagedsqlserverserverencryptionprotector

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceManagedSqlServerServerEncryptionProtectorClient struct {
	Client  autorest.Client
	baseUri string
}

func NewWorkspaceManagedSqlServerServerEncryptionProtectorClientWithBaseURI(endpoint string) WorkspaceManagedSqlServerServerEncryptionProtectorClient {
	return WorkspaceManagedSqlServerServerEncryptionProtectorClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
