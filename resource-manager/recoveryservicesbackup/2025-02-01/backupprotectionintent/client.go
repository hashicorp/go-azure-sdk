package backupprotectionintent

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackupProtectionIntentClient struct {
	Client  autorest.Client
	baseUri string
}

func NewBackupProtectionIntentClientWithBaseURI(endpoint string) BackupProtectionIntentClient {
	return BackupProtectionIntentClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
