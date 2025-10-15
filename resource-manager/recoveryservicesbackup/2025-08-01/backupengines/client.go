package backupengines

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackupEnginesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewBackupEnginesClientWithBaseURI(endpoint string) BackupEnginesClient {
	return BackupEnginesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
