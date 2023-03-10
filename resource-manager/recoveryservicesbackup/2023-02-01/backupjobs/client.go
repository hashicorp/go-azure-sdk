package backupjobs

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackupJobsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewBackupJobsClientWithBaseURI(endpoint string) BackupJobsClient {
	return BackupJobsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
